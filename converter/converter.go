/**
 *Author: Stefan
 *Date: 11/27/2019
 *Last changes: 11/30/2019 20.20
 *Task: Implement string to int converter, like
		● func myStrToInt(s str) (int, error){}
		● Cover it with tests
*/

package converter

import "strconv"

//MyStrToInt func
func MyStrToInt(s string) (int, error) {
	if v, err := strconv.Atoi(s); err == nil {
		return v, nil
	} else {
		return 0, err
	}
}
