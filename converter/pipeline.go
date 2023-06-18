package converter

import (
	"bufio"
	"context"
	"github.com/metapox/sayu/interface"
	"os"
	"sync"
)

type pipeline struct {
	input      _interface.Input
	converters []_interface.Converter
	workers    []Worker
	wg         sync.WaitGroup
	queue      chan []byte
}

func (pipeline *pipeline) Start() (err error) {
	inp, err := os.Open("test/test.log")
	defer inp.Close()
	if err != nil {
		return err
	}

	fr := bufio.NewScanner(inp)

	for fr.Scan() {
		go pipeline.workers[0].Push(fr.Bytes())
	}
	ctx, cancel := context.WithCancel(context.Background())
	nctx, ncancel := context.WithCancel(context.Background())
	cancel()
	for _, worker := range pipeline.workers {
		pipeline.wg.Add(1)
		go func() {
			go worker.Start(ctx)
			worker.Wait()
			pipeline.wg.Done()
			ncancel()
		}()
		ctx = nctx
		nctx, ncancel = context.WithCancel(context.Background())
	}
	out, err := os.OpenFile("test/new_test.log", os.O_WRONLY|os.O_CREATE, 0666)
	defer out.Close()
	go func() {
		for data := range pipeline.queue {
			out.Write(data)
			out.Write([]byte("\n"))
		}
	}()
	pipeline.wg.Wait()

	return nil
}

func (pipeline *pipeline) ShowConvertersInfo() string {
	info := "以下のconverterが登録されています\n"
	for _, converter := range pipeline.converters {
		info += converter.Name() + "\n"
	}
	return info
}

func (pipeline *pipeline) RegistConverter(converter _interface.Converter) {
	nextQueue := make(chan []byte, 1)
	pipeline.workers = append(pipeline.workers, NewWorker(100, converter, pipeline.queue, nextQueue))
	pipeline.queue = nextQueue
	pipeline.wg.Add(1)
}

func NewPipeline(input _interface.Input) *pipeline {
	return &pipeline{
		input:      input,
		converters: []_interface.Converter{},
		workers:    []Worker{},
		queue:      make(chan []byte, 1),
	}
}
