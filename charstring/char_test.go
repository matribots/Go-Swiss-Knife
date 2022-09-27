package charstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomPrintableASCII(t *testing.T) {
	assert.PanicsWithValue(t, "The \"len\" must gt 0", func() { RandomPrintableASCII(33, 88, -1) })
	assert.PanicsWithValue(t, "The \"min\" must lt \"max\"", func() { RandomPrintableASCII(33, 33, 20) })
	assert.PanicsWithValue(t, "The \"min\" must lt \"max\"", func() { RandomPrintableASCII(100, 33, 20) })
	assert.PanicsWithValue(t, "The \"min\" must ge 33", func() { RandomPrintableASCII(10, 200, 20) })
	assert.PanicsWithValue(t, "The \"max\" must lt 127", func() { RandomPrintableASCII(35, 130, 20) })
	asciiChars := RandomPrintableASCII(33, 127, 20)
	for _, asciiChar := range asciiChars {
		t.Log(string(asciiChar))
	}
}
