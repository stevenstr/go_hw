/*
 *Author: Stefan
 *Date: 11/24/2019
 *Last changes: 11/26/2019 13.20
 *Task: Implement a HelloWorld program that will print a smiley us
party dependency (https://github.com/kyokomi/emoji).
**/

package main

import (
	"fmt"

	"github.com/kyokomi/emoji"
)

func main() {
	_, err := emoji.Println(" :beer: :smile: :beer: :smile: :beer: :smile:")
	if err != nil {
		fmt.Println(err)
	}
}
