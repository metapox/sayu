package _interface

type Input interface {
	Hash() []byte
	Read(chan<- []byte)
}
