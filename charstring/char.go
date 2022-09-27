package charstring

import "github.com/matribots/go-swiss-knife/numbers"

type ascii byte

// Generate printable ASCII characters with given length that locate 32 < [min,max) < 127 safely.
// For example: RandomPrintableASCII(58, 100, 20) will generate 20 ASCII characters that locate in [58, 100)
func RandomPrintableASCII(min, max, len int) []ascii {
	if len <= 0 {
		panic("The \"len\" must gt 0")
	}
	if min >= max {
		panic("The \"min\" must lt \"max\"")
	}
	if min < 33 {
		panic("The \"min\" must ge 33")
	}
	if max > 127 {
		panic("The \"max\" must lt 127")
	}

	asciiChar := make([]ascii, len)

	for i := 0; i < len; i++ {
		n := numbers.RandomPostiveInt(int64(max))
		// Regenerate if generated num is smaller than min
		for n < min {
			n = numbers.RandomPostiveInt(int64(max))
		}
		asciiChar[i] = ascii(n)
	}

	return asciiChar
}
