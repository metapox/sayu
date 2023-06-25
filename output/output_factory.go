package output

import (
	"fmt"
	mconfig "github.com/metapox/sayu/config"
	_interface "github.com/metapox/sayu/interface"
	"gopkg.in/yaml.v3"
)

func CreateOutput(conf mconfig.Output) (_interface.Output, error) {
	var i _interface.Output
	var err error
	opts, _ := yaml.Marshal(conf.Options)

	switch conf.Name {
	case "localfile":
		o := LocalfileOptions{}
		err = yaml.Unmarshal(opts, &o)
		if err != nil {
			return nil, fmt.Errorf("output: invalid options")
		}
		i, err = NewLocalfile(o)
	default:
		return nil, fmt.Errorf("output: %s not found", conf.Name)
	}
	return i, err
}
