package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// Brute force
func longestPalindromeBruteForce(s string) string {
	strLen := len(s)
	var maxLen, start int
	for i := 0; i < strLen; i++ {
		for j := i; j < strLen; j++ {
			flag := 1

			for k := 0; k < ((j-i)/2)+1; k++ {
				if s[i+k] != s[j-k] {
					flag = 0
				}
			}

			if flag != 0 && (j-i+1) > maxLen {
				start = i
				maxLen = j - i + 1
			}
		}
	}
	return s[start : start+maxLen]
}

// Expand around center
func longestPalindrome(s string) string {
	strLen := len(s)
	if strLen == 0 {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < strLen; i++ {
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)
		len := int(math.Max(float64(len1), float64(len2)))
		if len > end-start {
			start = i - (len-1)/2
			end = i + (len)/2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left int, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}

func main() {
	r := bufio.NewReaderSize(os.Stdin, 1000)
	input, _ := r.ReadString('\n')
	//out := longestPalindrome(input)
	out := longestPalindromeBruteForce(input)
	fmt.Println(out)
}
