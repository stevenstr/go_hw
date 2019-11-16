/*
 *Author: Stefan
 *Date: 11/13/2019
 *Task: Implement the following methods: End - defines the end of the shape, Perimiter, Area
**/

package main

import "fmt"

//Point structure
type Point struct {
	x, y int
}

//Square structure
type Square struct {
	start Point
	a     uint
}

//End method
func (p *Square) End() (int, int) {
	var xend = p.start.x + int(p.a)
	var yend = p.start.y - int(p.a)
	return xend, yend
}

//Perimeter method
func (p *Square) Perimeter() uint {
	return p.a * uint(4)
}

//Area method
func (p *Square) Area() uint {
	return (p.a * p.a)
}

func main() {
	s := Square{Point{1, 1}, 5}
	fmt.Println(s.End())
	fmt.Println(s.Perimeter())
	fmt.Println(s.Area())
}
