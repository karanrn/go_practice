package main

import (
	"fmt"

	dist "github.com/karanrn/go_practice/distconv"
)

func main() {
	kms := dist.Kilometer(100)
	miles := dist.Mile(55)
	fmt.Printf("Kms: %.2f in Miles: %.2f \n", kms, dist.KmsToMiles(kms))
	fmt.Printf("Miles: %.2f in Kms: %.2f \n", miles, dist.MilesToKms(miles))
}
