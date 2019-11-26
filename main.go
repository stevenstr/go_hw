/*
 *Author: Stefan
 *Date: 11/24/2019
 *Last changes: 11/26/2019 11.10
 *Task: Implement Square and Circle structures which implements
		Figure interface.
		Add test cases for Perimeter / Area calculation.
**/

package main

import (
	"fmt"
	"math"
)

//Square struct
type Square struct {
	a float64
}

//Circle Struct
type Circle struct {
	radius float64
}

//area method for Circle
func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

//perimeter method for Circle
func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

//perimeter method for Square
func (p Square) perimeter() float64 {
	return p.a * 4
}

//area method for Square
func (p Square) area() float64 {
	return math.Pow(p.a, 2)
}

//Figure interface
type Figure interface {
	area() float64
	perimeter() float64
}

func main() {
	//some test data
	var s Figure = Square{5}
	var c Figure = Circle{3}
	fmt.Println(s.area(), s.perimeter())
	fmt.Println(c.area(), c.perimeter())

	s = Square{7}
	c = Circle{2}
	fmt.Println(s.area(), s.perimeter())
	fmt.Println(c.area(), c.perimeter())
}
