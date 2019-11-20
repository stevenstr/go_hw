/*
 *Author: Stefan
 *Date: 11/20/2019
 *Last changes: 11/20/2019 15.05
 *Task: Implement printSorted(map[int]string) function that prints map values sorted in order of increasing keys.
 *					Input -> {2: "a", 0: "b", 1: "c"}
 *					Output -> ["b", "c", "a"]
**/

package main

import (
	"fmt"
	"sort"
)

//printSorted function
func printSorted(m map[int]string) {

	var sl []int

	for i := range m {
		sl = append(sl, i)
	}
	//Is the easiest way to do it
	sort.Ints(sl)

	for _, i := range sl {
		fmt.Print(m[i], " ")
	}
	fmt.Println()
}

func main() {
	m := map[int]string{2: "a", 0: "b", 1: "c", 5: "g"}
	printSorted(m)
	m = map[int]string{10: "aa", 0: "bb", 500: "cc"}
	printSorted(m)
}
