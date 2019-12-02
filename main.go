/**
 *Author: Stefan
 *Date: 12/02/2019
 *Last changes: 12/02/2019 19.20
 *Task: Get the code from https://pastebin.com/9HCGfz26
 *		● Run go vet, fix errors.
 *		● Repeat until all is fixed.
 */

package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

//multiplyByTwo function
func multiplyByTwo(k *int) error {
	*k *= 2
	return nil
}

//printMoreTen function
func printMoreTen(g int) error {
	if g < 10 {
		return errors.New("g must be > 10")
	}
	fmt.Println(g)
	return nil
}

//jsStruct structure
type jsStruct struct {
	Data int  `json:"data"`
	OK   bool `json:"ok"`
}

//dejson function
func dejson() (jsStruct, error) {
	in := []byte(`{"data": 13, "ok": true}`)
	var out jsStruct
	if err := json.Unmarshal(in, &out); err != nil { //https://golang.org/pkg/encoding/json/#Unmarshal
		panic(err)
	}
	return out, nil
}

func main() {
	var r int = 11
	multiplyByTwo(&r)
	err := printMoreTen(r)
	if err != nil {
		panic(err)
	}
}
