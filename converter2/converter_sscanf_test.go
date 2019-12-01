/**
 *Author: Stefan
 *Date: 11/27/2019
 *Last changes: 12/01/2019 20.20
 *Task: Tests and benchmark for Converster which using fmt Sscanf
 */

package converter2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//TestDoneConvert function
func TestDoneConvert(t *testing.T) {
	result, err := MyStrToIntSscanf("1234")

	assert.Equal(t, result, 1234, "should be equal")
	assert.Nil(t, err, "should be nil")
}

//TestMixedConvert function, if the string consist of symbols and numbers
func TestMixedConvert(t *testing.T) {
	r, err := MyStrToIntSscanf("hhhhh1g124ab666c")

	assert.Equal(t, r, 0, "should be equal")
	assert.Error(t, err, "should be error")
}

//TestEmptyConvert function, if the string an empty
func TestEmptyConvert(t *testing.T) {
	result, err := MyStrToIntSscanf("")

	assert.Equal(t, result, 0, "should be equal")
	assert.Error(t, err, "should be error")
}

//TestOverBuffConvert function, if the stryng can't be coverted
func TestOverBuffConvert(t *testing.T) {
	result, err := MyStrToIntSscanf("999999999999999999999999999999999999999999")

	assert.Equal(t, result, 0, "should be equal")
	assert.Error(t, err, "should be error")
}

//BenchmarkMyStrToIntSscanf function
func BenchmarkMyStrToIntSscanf(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		_, _ = MyStrToIntSscanf(string(i))
	}
}
