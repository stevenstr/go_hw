/**
 *Author: Stefan
 *Date: 11/27/2019
 *Last changes: 12/01/2019 20.20
 *Task: Implement string to int converter, like
		● func myStrToInt(s str) (int, error){}
*/

package converter

import "strconv"

//MyStrToIntAtoi func
func MyStrToIntAtoi(s string) (int, error) {
	if v, err := strconv.Atoi(s); err == nil {
		return v, nil
	} else {
		return 0, err
	}
}
