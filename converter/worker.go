package converter

import (
	"context"
	_interface "github.com/metapox/sayu/interface"
	"log"
	"sync"
)

type Worker struct {
	sem       chan struct{}
	queue     chan []byte
	nextQueue chan []byte
	data      []byte
	converter _interface.Converter
	wg        sync.WaitGroup
}

func NewWorker(maxConcurrentExecution int, converter _interface.Converter, queue chan []byte, nextQueue chan []byte) Worker {
	return Worker{
		sem:       make(chan struct{}, maxConcurrentExecution),
		queue:     queue,
		nextQueue: nextQueue,
		data:      []byte{},
		converter: converter,
		wg:        sync.WaitGroup{},
	}
}

func (worker Worker) Start(ctx context.Context) {
	worker.wg.Add(1)
	go worker.loop(ctx)
}

func (worker Worker) Push(data []byte) {
	worker.queue <- data
}

func (worker Worker) Stop() {
	worker.wg.Done()
}

func (worker Worker) Wait() {
	worker.wg.Wait()
}

func (worker Worker) loop(ctx context.Context) {
	var wg sync.WaitGroup
Loop:
	for {
		select {
		case <-ctx.Done():
			log.Println(worker.converter.Name(), "start")
			wg.Wait()
			break Loop
		case data := <-worker.queue:
			wg.Add(1)
			worker.sem <- struct{}{}
			go func() {
				defer func() {
					<-worker.sem
					wg.Done()
				}()
				var ndata []byte
				ndata, worker.data = worker.converter.Convert(worker.data, data)
				if ndata != nil {
					worker.nextQueue <- ndata
				}
			}()
		}
	}
	worker.Stop()
}
