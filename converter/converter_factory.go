package converter

import (
	"fmt"
	mconfig "github.com/metapox/sayu/config"
	"github.com/metapox/sayu/converter/converters"
	_interface "github.com/metapox/sayu/interface"
)

func CreateConverter(conf mconfig.Converter) (_interface.Converter, error) {
	var c _interface.Converter
	var err error

	switch conf.Name {
	case "logConverter":
		c, err = converters.NewLogConverter()
	default:
		return nil, fmt.Errorf("%s not found", conf.Name)
	}
	return c, err
}
