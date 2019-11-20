/*
 *Author: Stefan
 *Date: 11/20/2019
 *Last changes: 11/20/2019 12.50
 *Task: Implement average([]int) float64 function that returns an average value of slice (sum / N)
**/

package main

import "fmt"

func average(sl []int) float64 {
	var N int = len(sl)
	var sum int = 0
	for i := 0; i < N; i++ {
		sum += sl[i]
	}
	return float64(sum) / float64(N)
}

func main() {
	var sl = []int{1, 2, 3, 4, 5, 6}
	fmt.Println(average(sl))
}
