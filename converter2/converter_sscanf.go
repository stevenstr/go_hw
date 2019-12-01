/**
 *Author: Stefan
 *Date: 11/27/2019
 *Last changes: 11/30/2019 20.20
 *Task: Implement string to int converter, like
		● func myStrToInt(s str) (int, error){}
		● Cover it with tests
*/

package converter2

import "fmt"

//MyStrToIntSscanf func
func MyStrToIntSscanf(s string) (int, error) {
	var a int
	if v, err := fmt.Sscanf(s, "%d", &a); err == nil {
		return v, nil
	} else {
		return 0, err
	}
}
