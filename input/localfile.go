package input

import (
	"bufio"
	"crypto/sha256"
	"io"
	"log"
	"os"
)

type Localfile struct {
	hash    []byte
	options LocalfileOptions
}

type LocalfileOptions struct {
	Filepath string `yaml:"filepath"`
}

func NewLocalfile(options LocalfileOptions) (*Localfile, error) {
	file, err := os.Open(options.Filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		log.Fatal(err)
	}

	hashValue := hash.Sum(nil)

	return &Localfile{
		options: options,
		hash:    hashValue,
	}, err
}

func (localfile *Localfile) Read(queue chan<- []byte) {
	file, err := os.Open(localfile.options.Filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fr := bufio.NewScanner(file)
	for fr.Scan() {
		queue <- fr.Bytes()
	}
}

func (localfile *Localfile) Hash() []byte {
	return localfile.hash
}
