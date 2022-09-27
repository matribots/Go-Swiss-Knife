package charstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEmptyString(t *testing.T) {
	var str1 *string
	var str2 *string = nil
	var str3 string = ""
	var str4 string = "     "
	var str5 string = "NOT_EMPTY"
	var str6 string = "NOT E MPTY"
	assert.PanicsWithValue(t, "The param \"str\" cannot be nil!", func() { IsEmptyString(str1) })
	assert.PanicsWithValue(t, "The param \"str\" cannot be nil!", func() { IsEmptyString(str2) })
	assert.True(t, IsEmptyString(&str3))
	assert.True(t, IsEmptyString(&str4))
	assert.False(t, IsEmptyString(&str5))
	assert.False(t, IsEmptyString(&str6))
}
