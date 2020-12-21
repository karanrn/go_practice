package main

import (
	"fmt"
	"strings"
)

var roman = map[int]string{
	1:    "I",
	5:    "V",
	10:   "X",
	50:   "L",
	100:  "C",
	500:  "D",
	1000: "M",
}

func getNum(digit int, letters []string) string {
	if digit > 0 && digit < 5 {
		if digit == 4 {
			return letters[0] + letters[1]
		}
		return strings.Repeat(letters[0], digit)
	} else if digit == 5 {
		return letters[1]
	} else if digit > 5 && digit < 9 {
		return letters[1] + strings.Repeat(letters[0], digit-5)
	} else if digit == 9 {
		return letters[0] + letters[2]
	}
	return ""
}

func intToRoman(num int) string {
	base := 1
	sol := ""
	for num > 0 {
		digit := num % 10
		num = num / 10

		switch base {
		case 1:
			sol = getNum(digit, []string{"I", "V", "X"}) + sol
		case 10:
			sol = getNum(digit, []string{"X", "L", "C"}) + sol
		case 100:
			sol = getNum(digit, []string{"C", "D", "M"}) + sol
		case 1000:
			sol = getNum(digit, []string{"M", " ", " "}) + sol
		}

		base = base * 10
	}
	return sol
	/*
		M := []string{"", "M", "MM", "MMM"}
		C := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
		X := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
		I := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}

		return M[num/1000] + C[(num%1000)/100] + X[(num%100)/10] + I[num%10]
	*/
}

func main() {
	var input int
	fmt.Print("Enter integer between 1 to 3999:")
	fmt.Scanln(&input)
	fmt.Println(intToRoman(input))
}
