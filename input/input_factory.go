package input

import (
	"fmt"
	mconfig "github.com/metapox/sayu/config"
	_interface "github.com/metapox/sayu/interface"
	"gopkg.in/yaml.v3"
)

func CreateInput(conf mconfig.Input) (_interface.Input, error) {
	var i _interface.Input
	var err error
	opts, _ := yaml.Marshal(conf.Options)

	switch conf.Name {
	case "localfile":
		o := LocalfileOptions{}
		err = yaml.Unmarshal(opts, &o)
		if err != nil {
			return nil, fmt.Errorf("input: invalid options")
		}
		i, err = NewLocalfile(o)
	default:
		return nil, fmt.Errorf("input: %s not found", conf.Name)
	}
	return i, err
}
