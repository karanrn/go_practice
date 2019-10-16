package main

import (
	"fmt"
	"sync"
	_ "time"
)

func consoleLog(chk int, c chan string, wg *sync.WaitGroup) {
	// Writing to channel to get the response.
	switch {
	case chk == 1:
		fmt.Println("You secured first place.")
		c <- "You secured first place."
	case chk == 2:
		fmt.Println("You secured second place.")
		c <- "You secured second place."
	case chk == 3:
		fmt.Println("You secured third place.")
		c <- "You secured third place."
	case 4 <= chk && chk <= 10:
		fmt.Println("You will be awarded consolation prize.")
		c <- "You will be awarded consolation prize."
	default:
		fmt.Println("Congrats for participating")
		c <- "Congrats for participating."
	}
	wg.Done()
}

func main() {
	fmt.Println("Sample program to test go routines.")
	// Create channel
	var ch chan string = make(chan string)

	// Wait groups for goroutines
	var waitgroup sync.WaitGroup
	// first routine
	waitgroup.Add(2)
	go consoleLog(1, ch, &waitgroup)
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
	go consoleLog(12, ch, &waitgroup)
	fmt.Println("\nDone #2")

	// Receive from channel
	first, last := <-ch, <-ch
	// We need to wait for reading from channel
	waitgroup.Wait()
	fmt.Printf("#1: %s \n#2: %s\n", first, last)
	fmt.Println("Completed execution.")
}
