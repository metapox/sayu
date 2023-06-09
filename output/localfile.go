package output

import (
	"crypto/sha256"
	"io"
	"log"
	"os"
)

type Localfile struct {
	options LocalfileOptions
	hash    []byte
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
	}, nil
}

func (localfile *Localfile) Write(queue <-chan []byte) {
	out, err := os.OpenFile(localfile.options.Filepath, os.O_WRONLY|os.O_CREATE, 0666)
	defer out.Close()
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	for data := range queue {
		out.Write(data)
		out.Write([]byte("\n"))
	}
}

func (localfile *Localfile) Hash() []byte {
	return localfile.hash
}
