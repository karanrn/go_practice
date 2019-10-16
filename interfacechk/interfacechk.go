package main

import (
	"fmt"
	"strings"
	"os"
	"bufio"
	"flag"
)

type change interface {
	Randomize() string
}

// Different types
type upper struct{
	input string
}

type lower struct{
	input string
}

func (s lower) Randomize() string{
	return strings.ToLower(s.input)
}

func (s upper) Randomize() string{
	return strings.ToUpper(s.input)
}

func main(){
	// read a string
	var str, output string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the input string: ")
	str, _ = reader.ReadString('\n')

	// Command line flag parsing
	lowercase := flag.Bool("lower", true, "Convert to lower case.")
	uppercase := flag.Bool("upper", false, "Convert to upper case.")

	flag.Parse()
	// Implement the interfaces
	if (*uppercase){
		var upperstr change = upper{str}
		// Randomize the input
		output = upperstr.Randomize()
		fmt.Print("Changed output from upper => Randomize(): " + output)
	}else if (*lowercase){
		var lowerstr change = lower{str}
		// Randomize the input
		output = lowerstr.Randomize()
		fmt.Print("Changed output from lower => Randomize(): " + output)
	}	
}