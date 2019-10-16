package stringutils

import "unicode"

func AlternateUpper(input string) string {
	var temp = make([]byte, len(input))
	for i :=0; i<len(input); i++{
		if i%2 == 0{
			temp[i] = byte(unicode.ToUpper(rune(input[i])))
		} else {
			temp[i] = byte(input[i])
		}
	}
	return string(temp)
}
