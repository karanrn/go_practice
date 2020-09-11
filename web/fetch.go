package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	// Get HTML source of the url passed
	for _, url := range os.Args[1:] {
		// Add http:// if not passed
		if !(strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")) {
			url = "http://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		// ioutil.ReadAll reads everything into memory
		/*
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Fprintf(os.Stderr, "fetch: reading %s: %v \n", url, err)
				os.Exit(1)
			}
			fmt.Printf("%s: %d", b, resp.StatusCode)
		*/

		// We shall use io.Copy to address it
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

	}
}
