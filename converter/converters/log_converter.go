package converters

import (
	"fmt"
	"regexp"
)

type LogConverter struct {
	name string
}

func NewLogConverter() (*LogConverter, error) {
	return &LogConverter{
		name: "log-converter",
	}, nil
}

func (converter LogConverter) Name() string {
	return converter.name
}

func (converter LogConverter) Convert(data []byte, ndata []byte) ([]byte, []byte) {
	data = append(data, ndata...)
	matched := regexp.MustCompile(`^(\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}\.\d{3}) (\w+) \(([\w-]+)\) \[([\w:]+)\] (.*)$`).Match(data)

	if matched {
		fmt.Println(string(data))
		return data, []byte{}
	} else {
		pattern := `\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}\.\d{3}`
		nmatched := regexp.MustCompile(pattern).Match(data)
		if nmatched {
			return nil, data
		} else {
			return nil, []byte{}
		}
	}
}
