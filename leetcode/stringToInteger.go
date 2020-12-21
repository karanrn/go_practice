package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isDigit(ch rune) bool {
	digits := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	for _, d := range digits {
		if ch == d {
			return true
		}
	}
	return false
}

// Function considers only the first occurence of number in the string
func myAtoi(s string) int {
	var intStr string
	input := strings.Trim(s, " ")
	var sign bool
	for i := 0; i < len(input); i++ { // _, s := range input {
		if isDigit(rune(input[i])) || (rune(input[i]) == '-' || rune(input[i]) == '+') && !sign && (i+1 < len(input) && isDigit(rune(input[i+1]))) {
			sign = true
			intStr = intStr + string(input[i])
		} else {
			break
		}
	}
	fmt.Println(intStr)
	out, _ := strconv.ParseInt(intStr, 10, 32)
	return int(out)
}

func main() {
	in := bufio.NewReaderSize(os.Stdin, 200)
	input, _ := in.ReadString('\n')
	out := myAtoi(input)
	fmt.Println(out)
}
