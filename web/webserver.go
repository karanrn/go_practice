// Minimal "echo" server
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"golang.org/x/time/rate"

	"github.com/dgrijalva/jwt-go"
)

// Implementing middleware for authentication
func authenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			fmt.Println("Malformed token")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Malformed Token"))
		} else {
			jwtToken := authHeader[1]

			token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}
				return []byte("ASFF!@#KJSKAJD"), nil
			})

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				ctx := context.WithValue(r.Context(), "props", claims)
				next.ServeHTTP(w, r.WithContext(ctx))
			} else {
				fmt.Println(err)
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized"))
			}
		}
	})
}

// logging middleware
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		file, err := os.OpenFile("access.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Fprintf(w, "failed to open file: %v", err)
			return
		}
		defer file.Close()

		fmt.Fprintf(file, "%v: You will be authenticated first.\n", time.Now())
		next.ServeHTTP(w, r)
		fmt.Fprintf(file, "%v: You are done.\n", time.Now())
	})
}

// limiters
var limiters = make(map[string]*rate.Limiter)

func init() {
	limiters["Authorized"] = rate.NewLimiter(2, 10)
	limiters["Unauthorized"] = rate.NewLimiter(1, 3)
}

// RateLimitMiddleware .
func RateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") == "JOHN_SNOW" {
			if !limiters["Authorized"].Allow() {
				http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
				return
			}
		} else {
			if !limiters["Unauthorized"].Allow() {
				http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", RateLimitMiddleware(mux)))
}

// handler echoes the path component
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
