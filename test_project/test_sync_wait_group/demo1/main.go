package main

import (
	"fmt"
	"sync"
	"time"
)

/**
sync.WaitGroup会阻塞等待所有任务完成之后再继续执行
*/

type httpPkg struct{}

func (httpPkg) Get(url string) {

	time.Sleep(1 * time.Second)
	fmt.Printf("finish to consume %s\n", url)
}

var http httpPkg

/**
sync.WaitGroup会阻塞等待所有任务完成之后再继续执行
*/
func main() {
	var wg sync.WaitGroup
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
	}
	for _, url := range urls {
		// Increment the WaitGroup counter.
		wg.Add(1)
		// Launch a goroutine to fetch the URL.
		go func(url string) {
			// Decrement the counter when the goroutine completes.
			defer wg.Done()
			// Fetch the URL.
			// time.Sleep(1 * time.Second)
			http.Get(url)
		}(url)
	}
	// Wait for all HTTP fetches to complete.
	wg.Wait()
	fmt.Print("all finish")
}
