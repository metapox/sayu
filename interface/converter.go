package _interface

type Converter interface {
	Name() string
	Convert([]byte, []byte) ([]byte, []byte)
}
