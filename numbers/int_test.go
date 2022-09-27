package numbers

import (
	"crypto/rand"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomPostiveInt(t *testing.T) {
	var ceil int64 = 1000
	testTimes := 100000
	invalidMax := [10]int64{0, -1, -2, -3, -4, -5, -6, -7, -8, -9}
	// The max value should locate in (0, ceil)
	autoGenValidMax := func(ceil int64) int64 {
		max, err := rand.Int(rand.Reader, big.NewInt(ceil))
		if err != nil {
			panic(err)
		}
		for int64(max.Uint64()) == 0 {
			max, err = rand.Int(rand.Reader, big.NewInt(ceil))
			if err != nil {
				panic(err)
			}
		}
		return int64(max.Uint64())
	}

	// Test invalid inputs
	for _, max := range invalidMax {
		assert.PanicsWithValue(t, "\"max\" must greater than 0", func() { RandomPostiveInt(max) })
	}

	// Test RandomPostiveInt(max int64) with given times
	for i := 0; i < testTimes; i++ {
		max := autoGenValidMax(ceil)
		randomNum := RandomPostiveInt(max)
		t.Logf("The random number is: %d; The max number is: %d", randomNum, max)
		// randomNum should located in [0,max]
		assert.GreaterOrEqual(t, randomNum, 0)
		assert.LessOrEqual(t, randomNum, int(max))
	}
}

func TestRandomPostiveInts(t *testing.T) {
	// Test invalid input first
	min, max, len := -1, 20, 80
	assert.PanicsWithValue(t, "The min num must ge 0", func() { RandomPostiveInts(min, max, len) })
	min, max, len = 0, 0, 80
	assert.PanicsWithValue(t, "The min num must lt max", func() { RandomPostiveInts(min, max, len) })
	min, max, len = 20, 20, 80
	assert.PanicsWithValue(t, "The min num must lt max", func() { RandomPostiveInts(min, max, len) })
	min, max, len = 30, 20, 80
	assert.PanicsWithValue(t, "The min num must lt max", func() { RandomPostiveInts(min, max, len) })

	// Test valid input
	min, max, len = 0, 127, 20
	nums := RandomPostiveInts(min, max, len)
	for i, n := range nums {
		t.Logf("The %d number of 20 is: %d", i+1, n)
		assert.Len(t, nums, len)
		assert.GreaterOrEqual(t, n, min)
		assert.Less(t, n, max)
	}
}
