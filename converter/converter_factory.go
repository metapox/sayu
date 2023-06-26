package converter

import (
	"fmt"
	mconfig "github.com/metapox/sayu/config"
	"github.com/metapox/sayu/converter/converters"
	_interface "github.com/metapox/sayu/interface"
	"gopkg.in/yaml.v3"
)

func CreateConverter(conf mconfig.Converter) (_interface.Converter, error) {
	var c _interface.Converter
	var err error
	opts, _ := yaml.Marshal(conf.Options)

	switch conf.Name {
	case "logConverter":
		o := converters.LogConverterOptions{}
		err = yaml.Unmarshal(opts, &o)
		if err != nil {
			return nil, fmt.Errorf("%s converter: invalid options", conf.Name)
		}
		c, err = converters.NewLogConverter(o)
	default:
		return nil, fmt.Errorf("%s not found", conf.Name)
	}
	return c, err
}
