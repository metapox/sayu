package _interface

type Converter interface {
	Name() string
	Convert(input Input) []byte
}
