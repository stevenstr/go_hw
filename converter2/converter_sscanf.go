/**
 *Author: Stefan
 *Date: 11/27/2019
 *Last changes: 12/01/2019 20.20
 *Task: Implement string to int converter, like
		● func myStrToInt(s str) (int, error){}
		● Using fmt.Sscanf
*/

package converter2

import "fmt"

//MyStrToIntSscanf func
func MyStrToIntSscanf(s string) (int, error) {
	var a int
	_, err := fmt.Sscanf(s, "%d", &a)
	if err == nil {
		return a, nil
	}
	return 0, err
}
