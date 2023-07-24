package converters

import (
	"regexp"
)

type AccessLogConverter struct {
	name    string
	options LogConverterOptions
}

type AccessLogConverterOptions struct {
	MaxConcurrent int `yaml:"maxConcurrent"`
}

func NewAccessLogConverter(options LogConverterOptions) (*AccessLogConverter, error) {
	return &AccessLogConverter{
		name:    "log-converter",
		options: options,
	}, nil
}

func (converter AccessLogConverter) Name() string {
	return converter.name
}

func (converter AccessLogConverter) MaxConcurrent() int {
	return converter.options.MaxConcurrent
}

func (converter AccessLogConverter) Convert(data []byte, ndata []byte) ([]byte, []byte) {
	const pattern = `^(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})\s-\s-\s\[(.+?)\]\s"(.+?)"\s(\d{3})\s(\d+?)\s"(.+?)"\s"(.+?)"\s"-(.*)"`
	matched := regexp.MustCompile(pattern).Match(ndata)

	if matched {
		return ndata, []byte{}
	} else {
		return nil, []byte{}
	}
}
