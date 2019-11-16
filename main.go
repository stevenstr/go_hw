/*
 *Author: Stefan
 *Date: 11/16/2019
 *Task: Home work 2.1 Implement - func factorial(i uint) uint - function without using recursion
 *Examples: Factorial(0) = 1; For enother: Factorial(3) = 1 * 2 * 3 = 6; Factorial(5) = 1 * 2 * 3 * 4 * 5 = 120;
**/

package main

import "fmt"

func factorial(i uint) uint {

	if i == 0 {
		return 1
	}

	var result uint = 1
	var a uint = 2

	for a <= i {
		result *= a
		a++
	}

	return result
}

func main() {
	fmt.Println(factorial(5))
}
