/**
 *Author: Stefan
 *Date: 11/24/2019
 *Last changes: 11/27/2019 17.50
 *Task: Implement sort.Interface interface for People type.
 *				Sort people in order of increasing age.
 *				If two people have the same age - sort them by name.
 *				Add test cases.
 */

package main

import (
	"fmt"
	"os"
	"sort"
	"time"
)

//Person struct
type Person struct {
	firstName string
	lastName  string
	birthDay  time.Time
}

//People struct
type People []Person

//Len method
func (p People) Len() int {
	return len(p)
}

//Less method
func (p People) Less(i, j int) bool {

	if p[i].birthDay.Sub(p[j].birthDay) > 0 {
		return true
	}
	if p[i].birthDay.Sub(p[j].birthDay) < 0 {
		return false
	}
	if p[i].firstName < p[j].firstName {
		return true
	}
	if p[i].firstName > p[j].firstName {
		return false
	}
	return p[i].lastName < p[j].lastName
}

//Swap method
func (p People) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	ivanIvanovDate1, err := time.Parse("2006-Jan-02", "2005-Aug-10")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	ivanIvanovDate2, err := time.Parse("2006-Jan-02", "2003-Aug-05")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	artiomIvanovDate, err := time.Parse("2006-Jan-02", "2005-Aug-10")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//create array of test data
	p := People{
		{"Ivan", "Ivanov", ivanIvanovDate1},
		{"Ivan", "Ivanov", ivanIvanovDate2},
		{"Artiom", "Ivanov", artiomIvanovDate},
	}

	sort.Sort(p)

	for _, q := range p {
		fmt.Println(q.firstName, q.lastName, q.birthDay)
	}
}
