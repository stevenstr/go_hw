/*
 *Author: Stefan
 *Date: 11/21/2019
 *Last changes: 11/20/2019 13.12
 *Task: Implement average([]int) float64 function that returns an average value of slice (sum / N)
**/

package main

import "fmt"

func average(sl []int) float64 {
	if len(sl) == 0 {
		fmt.Println("Empty slice")
		return 0
	}
	var N int = len(sl)
	var sum int //var sum int = 0
	for i := 0; i < N; i++ {
		sum += sl[i]
	}
	return float64(sum) / float64(N)
}

func main() {
	var sl = []int{1, 2, 3, 4, 5, 6}
	fmt.Println(average(sl))
	sl = []int{}
	fmt.Println(average(sl))
}
