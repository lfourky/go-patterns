package downloader

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Book struct {
	content []byte
}

type BookDownloader interface {
	Download(url string) (Book, error)
}

type downloader struct {
	pool waitingQueue
}

type waitingQueue chan struct{}

//New is used to acquire the implementation of the BookDownloader interface (which is the 'downloader' struct)
func New(poolSize int) BookDownloader {
	queue := make(waitingQueue, poolSize)
	go func() {
		for i := 0; i < poolSize; i++ {
			queue <- struct{}{}
		}
	}()
	return downloader{
		pool: queue,
	}
}

//Download returns imaginary (but empty) bytes we would get from an actual download
func (d downloader) Download(url string) (Book, error) {
	//Wait for an available worker to begin the download (blocking operation)
	<-d.pool
	defer func() {
		//Return the worker to the pool so another download can begin (unblocks the next channel receive)
		d.pool <- struct{}{}
	}()

	//It's our turn to download a book, so prepare for the download and return the downloaded data
	var b Book

	data, err := get(url)
	if err != nil {
		return b, err
	}

	b.content = data

	return b, nil
}

func get(url string) ([]byte, error) {
	fmt.Printf("downloading from url: %s\n", url)

	//Simulate actual download time by waiting for a random amount of time (but up to 3 seconds)
	time.Sleep(time.Second * time.Duration(rand.Intn(3)))

	//Random (occasional) error generator
	if shouldCauseError := rand.Intn(10) == 1; shouldCauseError {
		return []byte{}, errors.New("random error during download")
	}

	return []byte{}, nil
}
