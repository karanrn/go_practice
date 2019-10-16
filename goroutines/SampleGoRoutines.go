package main

import(
	"fmt"
	_ "time"
)

func consoleLog(chk int, c chan string) {
	// Writing to channel to get the response.
	switch{
	case chk == 1:
		c <- "You secured first place."
	case chk == 2:
		c <- "You secured second place."
	case chk == 3:
		c <- "You secured third place."
	case 4 <= chk && chk <= 10:
		c <- "You will awarded consolation prize."
	default:
		c <- "Congrats for participating."
	}
}

func main() {
	fmt.Println("Sample program to test go routines.")
	// Create channel
	var ch chan string = make(chan string)
	// first routine
	go consoleLog(1, ch)
	fmt.Println("Done #1")
	
	/*
	// progress status
	for i:=0;i<100;i++{
		for _,r := range `-\|/`{
			fmt.Printf("\r%d/100 - %c",i, r)
			time.Sleep(100 * time.Millisecond)
		}
	}
	*/
	// second routine
	go consoleLog(12, ch)
	fmt.Println("\nDone #2")

	// Receive from channel
	first, last := <-ch, <-ch
	fmt.Printf("#1: %s \n#2: %s\n", first, last)
	fmt.Println("Completed execution.")
}