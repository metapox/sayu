package converter

import (
	"bufio"
	"github.com/metapox/sayu/interface"
	"os"
)

type pipeline struct {
	input      _interface.Input
	converters []_interface.Converter
	workers    []Worker
}

func (pipeline *pipeline) Start() (err error) {
	inp, err := os.Open("test/test.log")
	defer inp.Close()
	if err != nil {
		return err
	}

	fr := bufio.NewScanner(inp)
	out, err := os.OpenFile("test/new_test.log", os.O_WRONLY|os.O_CREATE, 0666)
	defer out.Close()

	for _, worker := range pipeline.workers {
		for fr.Scan() {
			worker.Push(fr.Bytes())
			out.Write(worker.converter.Convert(fr.Bytes()))
			out.Write([]byte("\n"))
		}
	}

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
	pipeline.workers = append(pipeline.workers, NewWorker(converter))
}

func NewPipeline(input _interface.Input) *pipeline {
	return &pipeline{
		input: input,
	}
}
