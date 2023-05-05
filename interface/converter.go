package _interface

type Converter interface {
	Name() string
	Convert(data string) string
}
