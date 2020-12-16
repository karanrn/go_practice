package main

import (
	"fmt"
	"math"
	"strconv"
)

func reverse(x int) int {
	temp := strconv.Itoa(x)
	var isNegative bool
	if x < 0 {
		temp = temp[1:]
		isNegative = true
	}
	var resStr []rune
	if isNegative {
		resStr = append(resStr, '-')
	}
	for i := len(temp) - 1; i >= 0; i-- {
		resStr = append(resStr, rune(temp[i]))
	}

	res, _ := strconv.ParseInt(string(resStr), 10, 32)
	// Return zero if value is not in range of int32
	if int32(res) == int32(math.Pow(2, 31)-1) || int32(res) == -int32(math.Pow(2, 31)) {
		res = 0
	}
	return int(res)
}

func main() {
	var input int32
	fmt.Print("Enter a integer value:")
	fmt.Scanln(&input)
	res := reverse(int(input))
	fmt.Printf("Reveresed string: %d\n", res)
}
