package main

import (
	"fmt"
	"math"
)

// HANDS ON 1
// create a type square
// create a type circle
// attach a method to each that calculates area and returns it
// create a type shape which defines an interface as anything which has the area method
// create a func info which takes type shape and then prints the area
// create a value of type square
// create a value of type circle
// use func info to print the area of square
// use func info to print the area of circle

type square struct {
	side float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (s square) area() float64 {
	return s.side * s.side
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	sq1 := square{
		side: 3.4,
	}
	c1 := circle{
		radius: 2.5,
	}
	info(sq1)
	info(c1)
	// fmt.Printf("The area of square of side %v is %.2f.\n", sq1.side, info(sq1))
	// fmt.Printf("The area of circle of radius %v is %.2f.\n", c1.radius, info(c1))

}
