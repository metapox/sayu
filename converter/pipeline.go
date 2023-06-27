package converter

import (
	"github.com/metapox/sayu/interface"
	"sync"
)

type pipeline struct {
	input      _interface.Input
	output     _interface.Output
	converters []_interface.Converter
	workers    []Worker
	wg         sync.WaitGroup
	queue      chan []byte
}

func (pipeline *pipeline) Start() (err error) {
	pipeline.startupWorkers()

	pipeline.input.Read(pipeline.workers[0].queue)
	pipeline.workers[0].Close()

	pipeline.wg.Wait()
	pipeline.output.Write(pipeline.queue)

	return nil
}

func (pipeline *pipeline) ShowConvertersInfo() string {
	info := "以下のconverterが登録されています\n"
	for _, worker := range pipeline.workers {
		info += worker.converter.Name() + "\n"
	}
	return info
}

func (pipeline *pipeline) RegistConverter(converter _interface.Converter) {
	nextQueue := make(chan []byte, 1)
	pipeline.workers = append(pipeline.workers, NewWorker(100, converter, pipeline.queue, nextQueue))
	pipeline.queue = nextQueue
	pipeline.wg.Add(1)
}

func (pipeline *pipeline) startupWorkers() {
	for _, worker := range pipeline.workers {
		go func() {
			go worker.Start()
			worker.Wait()
			pipeline.wg.Done()
		}()
	}
}

func NewPipeline(input _interface.Input, output _interface.Output) *pipeline {
	return &pipeline{
		input:      input,
		output:     output,
		converters: []_interface.Converter{},
		workers:    []Worker{},
		queue:      make(chan []byte, 1),
	}
}
