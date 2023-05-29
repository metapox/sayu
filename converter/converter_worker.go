package converter

import (
	_interface "github.com/metapox/sayu/interface"
)

type Worker struct {
	queue     chan []byte
	data      []byte
	converter _interface.Converter
}

func (worker Worker) Load() {
	for ndata := range worker.queue {
		worker.converter.Convert(worker.data, ndata)
	}
}

func (worker Worker) Close() {
	close(worker.queue)
}
