/**
 *Author: Stefan
 *Date: 11/27/2019
 *Last changes: 12/03/2019 12.05
 *Task: Implement string to int converter, like
		● func myStrToInt(s str) (int, error){}
		● Cover it with tests
*/

package converter

import "strconv"

//MyStrToInt func
func MyStrToInt(s string) (int, error) {
	return strconv.Atoi(s)
}
