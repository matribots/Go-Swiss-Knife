package numbers

import (
	"crypto/rand"
	"math/big"
)

// Generate a random postive number that locate in [0,max) safely.
func RandomPostiveInt(max int64) int {
	if max <= 0 {
		panic("\"max\" must greater than 0")
	}
	n, e := rand.Int(rand.Reader, big.NewInt(max))
	if e != nil {
		panic(e)
	}
	return int(n.Uint64())
}

// Generate random postive numbers with given length that locate in [min,max) safely.
func RandomPostiveInts(min, max, len int) []int {
	if min < 0 {
		panic("The min num must ge 0")
	}

	if min >= max {
		panic("The min num must lt max")
	}

	nums := make([]int, len)

	for i := 0; i < len; i++ {
		n := RandomPostiveInt(int64(max))
		// Regenerate if the generated number is smaller than min
		for n < min {
			n = RandomPostiveInt(int64(max))
		}
		nums[i] = n
	}

	return nums
}
