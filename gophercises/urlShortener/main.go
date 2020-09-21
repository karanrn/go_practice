package main

import (
	"fmt"
	"net/http"

	"github.com/karanrn/go_practice/gophercises/urlShortener/urlshort"
)

func main() {
	mux := defaultMux()

	pathToURLs := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	//mapHandler := urlshort.NewHTTPRedirectHandler(urlshort.NewBaseURLMapper(pathToURLs), mux)
	mapHandler := urlshort.MapHandler(pathToURLs, mux)

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", mapHandler)

}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
