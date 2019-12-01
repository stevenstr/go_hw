/**
 *Author: Stefan
 *Date: 11/27/2019
 *Last changes: 11/30/2019 20.20
 *Task: Some tests for converter.go
 */

package converter2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//TestDoneConvert1 function
func TestDoneConvert1(t *testing.T) {
	result, err := MyStrToIntSscanf("1234")

	assert.Equal(t, result, 1234, "should be equal")
	assert.Nil(t, err, "should be nil")
}

//TestMixedConvert1 function, if the string consist of symbols and numbers
func TestMixedConvert1(t *testing.T) {
	result, err := MyStrToIntSscanf("12s34a")

	assert.Equal(t, result, 0, "should be equal")
	assert.Error(t, err, "should be error")
}

//TestEmptyConvert1 function, if the string an empty
func TestEmptyConvert1(t *testing.T) {
	result, err := MyStrToIntSscanf("")

	assert.Equal(t, result, 0, "should be equal")
	assert.Error(t, err, "should be error")
}

//TestOverBuffConvert1 function, if the stryng can't be coverted
func TestOverBuffConvert1(t *testing.T) {
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
