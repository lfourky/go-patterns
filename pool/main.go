package main

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/lfourky/go-patterns/pool/downloader"
)

const (
	booksAmount = 30 //Amount of books to download
	poolSize    = 5  //Amount of books being downloaded simultaneously
)

var wg sync.WaitGroup

func main() {
	urls := generateBookURLs(booksAmount, "localhost:8080")

	bookDownloader := downloader.New(poolSize)

	var successfulDownloads int32
	var failedDownloads int32

	for _, u := range urls {
		wg.Add(1)
		go func(bookURL string) {
			defer wg.Done()
			_, err := bookDownloader.Download(bookURL)
			if err != nil {
				atomic.AddInt32(&failedDownloads, 1)
				fmt.Printf("err downloading from url: %s, err: %s\n", bookURL, err.Error())
				return
			}
			atomic.AddInt32(&successfulDownloads, 1)
			fmt.Printf("successfuly downloaded from url: %s\n", bookURL)
		}(u)
	}
	wg.Wait()

	fmt.Printf("Done! Successful downloads: %d, failed downloads: %d\nTotal amount of books at the start: %d\n", successfulDownloads, failedDownloads, booksAmount)
}

//Used for generating mock data
func generateBookURLs(size int, baseURL string) []string {
	books := make([]string, size)

	for i := range books {
		books[i] = fmt.Sprintf("%s/%d", baseURL, i)
	}

	return books
}
