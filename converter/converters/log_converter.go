package converters

import (
	_interface "github.com/metapox/sayu/interface"
)

type LogConverter struct {
	name      string
	loadQueue [][]byte
	data      []byte
}

func NewLogConverter() *LogConverter {
	return &LogConverter{
		name: "log-converter",
	}
}

func (converter LogConverter) Name() string {
	return converter.name
}

func (converter LogConverter) Convert(input _interface.Input) []byte {
	return []byte("test" + string(input.Hash()))
}

func (converter LogConverter) Load() {
}
