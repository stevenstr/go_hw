/*
 *Author: Stefan
 *Date: 11/20/2019
 *Last changes: 11/21/2019 13.33
 *Task: Write reverse([]int64) []int64 function that returns the copy of the original slice in reverse order. The type of elements is int64.
 *						Input -> (1, 2, 5, 15)
 *						Output -> (15, 5, 2, 1)
**/

package main

import "fmt"

//reverse function
func reverse(sl []int64) []int64 {
	//check
	if len(sl) == 0 {
		fmt.Println("Empty slice")
		return nil
	}
	//Just make copy with original len and cap
	var cpSl = make([]int64, len(sl), cap(sl))

	var start int //var start = 0
	var end = len(sl) - 1
	//Just  do it
	for i := len(sl) - 1; i >= 0; i-- {
		cpSl[start] = sl[end]
		start++
		end--
	}
	return cpSl
}

func main() {
	var sl = []int64{1, 2, 5, 15, 20, 33, 1}

	fmt.Println("original: ", sl)
	fmt.Println()
	fmt.Println("reversed: ", reverse(sl))
	fmt.Println()
	sl = []int64{}
	fmt.Println("empty: ", reverse(sl))
}
