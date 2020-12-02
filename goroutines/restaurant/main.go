package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	/*
		Fixed number of tables - 50
		Fixed number of waiter/waitress - 20
		Fixed number of chefs - 10
	*/

	// Orders taken
	inOrder := make(chan int, 50)
	// Orders to be processed by chefs
	cookFood := make(chan int, 10)
	// Orders ready to be served
	outOrder := make(chan int, 50)

	wg.Add(150)
	for i := 1; i <= 100; i++ {
		inOrder <- i
		go takeOrder(inOrder, cookFood)
		go cookOrder(cookFood, outOrder)
		go serveOrder(outOrder)
	}
	wg.Wait()

}

func takeOrder(inOrder chan int, cookFood chan int) {
	defer wg.Done()
	table := <-inOrder
	fmt.Printf("Order for table %d is placed.\n", table)
	cookFood <- table
}

func cookOrder(cookFood chan int, outOrder chan int) {
	defer wg.Done()
	table := <-cookFood
	fmt.Printf("Order for table %d is cooking.\n", table)
	time.Sleep(time.Second * 2) // Cooking time
	outOrder <- table
}

func serveOrder(outOrder chan int) {
	defer wg.Done()
	table := <-outOrder
	fmt.Printf("Order for table %d is served.\n", table)
}
