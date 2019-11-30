package converter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//Done_Convert function
func Test_Done_Convert(t *testing.T) {
	result, err := MyStrToInt("1234")

	assert.Equal(t, result, 1234, "should be equal")
	assert.Nil(t, err, "should be nil")
}

//Mixed_Convert function
func Test_Mixed_Convert(t *testing.T) {
	result, err := MyStrToInt("12s34a")

	assert.Equal(t, result, 0, "should be equal")
	assert.Error(t, err, "should be error")
}

//Empty_Convert function
func Test_Empty_Convert(t *testing.T) {
	result, err := MyStrToInt("")

	assert.Equal(t, result, 0, "should be equal")
	assert.Error(t, err, "should be error")
}

//OverBuff_Convert function
func Test_OverBuff_Convert(t *testing.T) {
	result, err := MyStrToInt("999999999999999999999999999999999999")

	assert.Equal(t, result, 0, "should be equal")
	assert.Error(t, err, "should be error")
}
