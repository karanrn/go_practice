package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

// Customer defines details for the customer's purchase
type Customer struct {
	Name       string
	City       string
	DayOfMonth int32
	Item       string
	Price      float32
}

func main() {
	// Read CSV file
	csvFile, err := os.Open("Customers.csv")
	if err != nil {
		fmt.Printf("error while opening file: %v", err.Error())
		return
	}
	defer csvFile.Close()
	
	r := csv.NewReader(bufio.NewReader(csvFile))
	var customers []Customer

	// To ignore header record (assuming we always get header record)
	_, err = r.Read()
	if err == io.EOF {
		return
	} else if err != nil {
		fmt.Printf("error reading from csv: %v", err.Error())
	}

	for {
		line, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading from csv: %v", err.Error())
		}
		dom, err := strconv.ParseInt(line[2], 10, 32)
		if err != nil {
			fmt.Printf("error while onverting string to int: %v", err.Error())
		}
		// removing $ from the string value
		price, err := strconv.ParseFloat(line[4][1:], 32)
		if err != nil {
			fmt.Printf("error while onverting string to float: %v", err.Error())
		}
		customers = append(customers, Customer{
			Name:       line[0],
			City:       line[1],
			DayOfMonth: int32(dom),
			Item:       line[3],
			Price:      float32(price),
		})
	}

	// List of customers from Manhattan
	for _, cust := range customers {
		if cust.City == "Manhattan" {
			fmt.Println(cust.Name)
		}
	}
}
