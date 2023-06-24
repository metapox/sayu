package converter

import (
	_interface "github.com/metapox/sayu/interface"
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

func (worker Worker) Start() {
	worker.wg.Add(1)
	go worker.loop()
}

func (worker Worker) Push(data []byte) {
	worker.queue <- data
}

func (worker Worker) Close() {
	close(worker.queue)
}

func (worker Worker) Stop() {
	worker.wg.Done()
	close(worker.nextQueue)
}

func (worker Worker) Wait() {
	worker.wg.Wait()
}

func (worker Worker) loop() {
	var wg sync.WaitGroup
	for data := range worker.queue {
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
	worker.Stop()
}
