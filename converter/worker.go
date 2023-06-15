package converter

import (
	_interface "github.com/metapox/sayu/interface"
	"sync"
)

type Worker struct {
	sem       chan struct{}
	queue     chan []byte
	converter _interface.Converter
	wg        sync.WaitGroup
}

func NewWorker(maxConcurrentExecution int, buffers int, converter _interface.Converter) Worker {
	return Worker{
		sem:       make(chan struct{}, maxConcurrentExecution),
		queue:     make(chan []byte, buffers),
		converter: converter,
		wg:        sync.WaitGroup{},
	}
}

func (worker Worker) Start() {
}

func (worker Worker) loop() string {
}

func (worker Worker) Push(data []byte) {
}

func (worker Worker) Close() {
}
