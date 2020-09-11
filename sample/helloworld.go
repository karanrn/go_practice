package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	str "github.com/karanrn/go_practice/stringutils"
)

func main() {
	// Start of execution
	start := time.Now()

	// Seeding with time to get random number
	rand.Seed(time.Now().UnixNano())
	var people int = rand.Intn(1000)
	// Hello world (World in kannada), checking UTF-8
	fmt.Println("Hello ಪ್ರಪಂಚ")
	// Converting argument to alternative upper case string
	var owner string
	if len(os.Args) >= 2 {
		owner = str.AlternateUpper(os.Args[1])
	} else {
		owner = "God"
	}
	fmt.Printf("%s welcomes %d people to his world.", owner, people)
	// End of execution
	elapsed := time.Since(start)
	fmt.Printf("Program took %s\n", elapsed)
}
