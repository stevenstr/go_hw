/**
 *Author: Stefan
 *Date: 11/27/2019
 *Last changes: 11/30/2019 20.20
 *Task: Some tests for converter which use an Atoi. And benchmark
 */

package converter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//TestDoneConvert function
func TestDoneConvert(t *testing.T) {
	result, err := MyStrToIntAtoi("1234")

	assert.Equal(t, result, 1234, "should be equal")
	assert.Nil(t, err, "should be nil")
}

//TestMixedConvert function, if the string consist of symbols and numbers
func TestMixedConvert(t *testing.T) {
	result, err := MyStrToIntAtoi("j12s34a")

	assert.Equal(t, result, 0, "should be equal")
	assert.Error(t, err, "should be error")
}

//TestEmptyConvert function, if the string an empty
func TestEmptyConvert(t *testing.T) {
	result, err := MyStrToIntAtoi("")

	assert.Equal(t, result, 0, "should be equal")
	assert.Error(t, err, "should be error")
}

//TestOverBuffConvert function, if the stryng can't be coverted
func TestOverBuffConvert(t *testing.T) {
	result, err := MyStrToIntAtoi("99999999999999999999999999999999999999954564564569")

	assert.Equal(t, result, 9223372036854775807, "should be error")
	assert.Error(t, err, "should be error")
}

//BenchmarkMyStrToIntAtoi function
func BenchmarkMyStrToIntAtoi(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		_, _ = MyStrToIntAtoi(string(i))
	}
}
