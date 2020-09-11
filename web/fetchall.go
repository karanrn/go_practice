package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		// start a goroutine
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch) // Read from channel
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

// Fetch url and write to channel
func fetch(url string, ch chan<- string) {
	start := time.Now()

	if !(strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")) {
		url = "http://" + url
	}

	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // Send to channel
		return
	}

	// We will write ioutil.Discard for demonstration
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("While reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
