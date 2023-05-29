package converters

import (
	queue "github.com/metapox/sayu/internal/pkg"
	"regexp"
)

type LogConverter struct {
	name      string
	loadQueue queue.Queue
	loadData  []byte
	dataQueue queue.Queue
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
	nd := converter.loadQueue.Dequeue().([]byte)
	pattern := `\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}\.\d{3}`
	matched := regexp.MustCompile(pattern).Match(converter.loadData)
	nmatched := regexp.MustCompile(pattern).Match(nd)

	if matched && nmatched {
		return data, ndata
	} else {
		data = append(data, nd...)
		return nil, data
	}
}
