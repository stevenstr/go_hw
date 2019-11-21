/*
 *Author: Stefan
 *Date: 11/20/2019
 *Last changes: 11/21/2019 13.21
 *Task: Write max([]string) string function that returns the longest word from the slice of strings (the first if there are more than one).
 *			Input -> ("one", "two", "three")
 *			Output -> ("three")
**/

package main

import (
	"fmt"
	"unicode/utf8"
)

func max(sl []string) string {
	//var length int = 0 //var length int = 0
	var N = len(sl)
	var result string //var result string = ""

	// if there are several elements of the same length, then it will return the first
	for i := 0; i < N; i++ {
		if utf8.RuneCountInString(sl[i]) > utf8.RuneCountInString(result) {
			//length = len(sl[i])
			result = sl[i]
		}
	}
	return result
}

func main() {
	var sl = []string{"one", "two", "three", "lesha", "ivan"}
	fmt.Println(max(sl))
	sl = []string{"one", "two"}
	fmt.Println(max(sl))
}
