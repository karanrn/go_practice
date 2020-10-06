package main

import (
	"fmt"
	"math"
)

type triangle struct {
	sides [3]float64
}

type square struct {
	side float64
}

type rectangle struct {
	length  float64
	breadth float64
}

type shape interface {
	area() float64
	perimeter() float64
}

func (s square) area() float64 {
	return s.side * s.side
}

func (r rectangle) area() float64 {
	return r.length * r.breadth
}

func (t triangle) area() float64 {
	s := (t.sides[0] + t.sides[1] + t.sides[2]) / 2
	return math.Sqrt(s * (s - t.sides[0]) * (s - t.sides[1]) * (s - t.sides[2]))
}

func (s square) perimeter() float64 {
	return 4 * s.side
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.breadth + r.length)
}

func (t triangle) perimeter() float64 {
	return t.sides[0] + t.sides[1] + t.sides[2]
}

func main() {
	// Enter square sides
	var s square
	fmt.Print("Enter the side of a square: ")
	fmt.Scanln(&s.side)
	// Enter rectangle sides
	var r rectangle
	fmt.Print("Enter the length of a rectangle: ")
	fmt.Scanln(&r.length)
	fmt.Print("Enter the breadth of a rectangle: ")
	fmt.Scanln(&r.breadth)
	// Enter triangle sides
	var t triangle
	fmt.Print("Enter sides of a triangle (space separated): ")
	fmt.Scan(&t.sides[0])
	fmt.Scan(&t.sides[1])
	fmt.Scanln(&t.sides[2])

	// Results
	fmt.Printf("Square (%.2f) - Area: %.2f, Perimeter: %.2f\n", s.side, s.area(), s.perimeter())
	fmt.Printf("Rectangle (%.2f, %.2f) - Area: %.2f, Perimeter: %.2f\n", r.length, r.breadth, r.area(), r.perimeter())
	fmt.Printf("Triangle (%.2f, %.2f, %.2f) - Area: %.2f, Perimeter: %.2f\n", t.sides[0], t.sides[1], t.sides[2], t.area(), t.perimeter())
}
