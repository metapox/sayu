package converters

import (
	"regexp"
)

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

func (converter LogConverter) Convert(data []byte, ndata []byte) ([]byte, []byte) {
	pattern := `\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}\.\d{3}`
	matched := regexp.MustCompile(pattern).Match(data)
	nmatched := regexp.MustCompile(pattern).Match(ndata)

	if matched && nmatched {
		return data, ndata
	} else {
		data = append(data, ndata...)
		return nil, data
	}
}
