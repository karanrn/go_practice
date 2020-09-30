package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fs := flag.NewFlagSet("txn", flag.ContinueOnError) //flag.Bool("print", false, "Display Hello world!")
	user := fs.String("user", "", "Enter user name")
	age := fs.Int64("age", 0, "Enter your age")

	if len(os.Args[1:]) < 1 {
		fmt.Println("You must pass subcommand")
		return
	}

	fs.Parse(os.Args[2:])
	if os.Args[1] == "txn" {
		if *user != "" {
			if *age != 0 {
				fmt.Printf("Hello %s, you are %d years old.\n", *user, *age)
			} else {
				fmt.Printf("Hello %s, welcome to the world.\n", *user)
			}
		} else {
			fmt.Println("Hello World")
		}
	}
}
