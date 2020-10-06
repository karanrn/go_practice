package main

import (
	"fmt"
	"time"
)

func fib(n int) int {
	// recursive
	if n < 2 {
		return n
	}

	return fib(n-1) + fib(n-2)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func main() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n)
	fmt.Printf("\rFibonacci (%d) = %d\n", n, fibN)
}
