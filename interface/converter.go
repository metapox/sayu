package _interface

type Converter interface {
	Name() string
	MaxConcurrent() int
	Convert([]byte, []byte) ([]byte, []byte)
}
