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
