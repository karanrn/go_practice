package main

import (
	"bufio"
	"fmt"
	"os"
)

// Longest substring with non repeating characters
func lengthOfLongestSubstring(s string) int {
	strLen := len(s)
	var globalCount int
	for i := 0; i < strLen; i++ {
		substr := map[rune]int{}
		count := 0
		for j := i; j < strLen; j++ {
			if substr[rune(s[j])] == 0 {
				substr[rune(s[j])] = 1
				count++
			} else {
				break
			}
		}
		if count >= globalCount {
			globalCount = count
		}
	}
	return globalCount
}

func main() {
	r := bufio.NewReader(os.Stdin)
	input, _ := r.ReadString('\n')
	out := lengthOfLongestSubstring(input)
	fmt.Println(out)
}
