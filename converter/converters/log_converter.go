package converters

import "fmt"

type LogConverter struct {
	name string
}

func NewLogConverter() *LogConverter {
	return &LogConverter{
		name: "log-converter",
	}
}

func (converter LogConverter) Name() string {
	return converter.name
}

func (converter LogConverter) Convert(data string) string {
	fmt.Println("LogConverter: " + data)
	return data
}
