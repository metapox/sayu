package output

import (
	"crypto/sha256"
	"io"
	"log"
	"os"
)

type LocalFile struct {
	filepath string
	hash     []byte
}

func NewLocalFile(filepath string) *LocalFile {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		log.Fatal(err)
	}

	hashValue := hash.Sum(nil)

	return &LocalFile{
		filepath: filepath,
		hash:     hashValue,
	}
}

func (localFile *LocalFile) Write(queue <-chan []byte) {
	out, err := os.OpenFile(localFile.filepath, os.O_WRONLY|os.O_CREATE, 0666)
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

func (localFile *LocalFile) Hash() []byte {
	return localFile.hash
}
