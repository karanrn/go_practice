package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// Echo the arguments passed
	// 1. For loop
	start := time.Now()
	s, sep := "", ""
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Printf("For loop: %q.", s)
	elapsed := time.Since(start)
	fmt.Println(" Execution time: ", elapsed)

	// 2. Range
	start = time.Now()
	s, sep = "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Printf("Range: %q.", s)
	elapsed = time.Since(start)
	fmt.Println(" Execution time: ", elapsed)

	// 3. strings.Join method
	start = time.Now()
	fmt.Printf("Join: %q.", (strings.Join(os.Args[1:], " ")))
	elapsed = time.Since(start)
	fmt.Println(" Execution time: ", elapsed)
}
