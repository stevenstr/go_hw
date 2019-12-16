/**
 *Author: Stefan
 *Date: 12/02/2019
 *Last changes: 12/16/2019 21.25
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
func multiplyByTwo(k *int) {
	*k *= 2
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
		return jsStruct{}, err
	}
	return out, nil
}

func main() {
	var r int = 11
	multiplyByTwo(&r)

	if err := printMoreTen(r); err != nil {
		panic(err)
	}
}
