package input

import (
	"bufio"
	"crypto/sha256"
	"io"
	"log"
	"os"
)

type SolrLog struct {
	filepath string
	hash     []byte
}

func NewSolrLog(filepath string) *SolrLog {
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

	return &SolrLog{
		filepath: filepath,
		hash:     hashValue,
	}
}

func (solrLog *SolrLog) Read(queue chan<- []byte) {
	file, err := os.Open(solrLog.filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fr := bufio.NewScanner(file)
	for fr.Scan() {
		queue <- fr.Bytes()
	}
}

func (solrLog *SolrLog) Hash() []byte {
	return solrLog.hash
}
