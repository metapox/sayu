package converter

import (
	"bufio"
	"fmt"
	"github.com/metapox/sayu/interface"
	"os"
)

type pipeline struct {
	converters []_interface.Converter
}

func (pipeline *pipeline) Start() (err error) {
	f, err := os.Open("test/test.log")
	defer f.Close()
	if err != nil {
		return err
	}

	fr := bufio.NewScanner(f)

	for fr.Scan() {
		fmt.Println(fr.Text())
		for _, converter := range pipeline.converters {
			converter.Convert(fr.Text())
		}
	}
	return nil
}

func (pipeline *pipeline) AddConverter(converter _interface.Converter) {
	pipeline.converters = append(pipeline.converters, converter)
}

func NewPipeline() *pipeline {
	return &pipeline{}
}
