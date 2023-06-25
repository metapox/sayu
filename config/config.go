package config

type Config struct {
	Input      Input       `yaml:"input"`
	Converters []Converter `yaml:"converters"`
	Output     Output      `yaml:"output"`
}

type Input struct {
	Name    string      `yaml:"name"`
	Options interface{} `yaml:"options"`
}

type Converter struct {
	Name    string      `yaml:"name"`
	Options interface{} `yaml:"options"`
}

type Output struct {
	Name    string      `yaml:"name"`
	Options interface{} `yaml:"options"`
}
