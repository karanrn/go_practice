package main

import (
	_ "bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Program to find duplicate lines from the files and count them
func main() {
	// We will use maps to store information of lines
	counts := make(map[string]map[string]int)

	for _, arg := range os.Args[1:] {
		counts[arg] = make(map[string]int)

		// Readfile returns byte[]
		data, err := ioutil.ReadFile(arg)
		if err != nil {
			fmt.Println("File reading error", err)
		}

		for _, line := range strings.Split(string(data), "\n") {
			counts[arg][line]++
		}
	}

	for file, lines := range counts {
		for line, n := range lines {
			if n > 1 {
				fmt.Printf("%d\t%s\t%s\n", n, line, file)
			}
		}
	}

}
