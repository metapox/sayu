package _interface

type Output interface {
	Hash() []byte
	Write(<-chan []byte)
}
