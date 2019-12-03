/**
 *Author: Stefan
 *Date: 11/27/2019
 *Last changes: 12/03/2019 12.10
 *Task: Implement string to int converter, like
		● func myStrToInt(s str) (int, error){}
*/

package converter

import "strconv"

//MyStrToIntAtoi func
func MyStrToIntAtoi(s string) (int, error) {
	return strconv.Atoi(s)
}
