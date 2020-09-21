package urlshort

import (
	"fmt"
	"net/http"
)

// MapHandler maps keys to URLs, redirects if path exists in the map
func MapHandler(pathToURLs map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if url, ok := pathToURLs[path]; ok {
			http.Redirect(w, r, url, http.StatusMovedPermanently)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

// NewBaseURLMapper maps keys to URLs, returns true if url exists
func NewBaseURLMapper(urls map[string]string) func(string) (string, bool) {
	return func(path string) (string, bool) {
		url, ok := urls[path]
		return url, ok
	}
}

// NewHTTPRedirectHandler redirects to URL if exists else fallbacks to http.Handler
func NewHTTPRedirectHandler(mapper func(string) (string, bool), fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if url, ok := mapper(r.URL.Path); ok {
			fmt.Printf("Redirecting %s to %s \n", r.URL.Path, url)
			http.Redirect(w, r, url, http.StatusMovedPermanently)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}

// YAMLHandler uses yaml mapping for URL redirection, if path does not exist it fallbacks to MapHandler
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	return nil, nil
}
