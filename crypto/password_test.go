package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidPassword(t *testing.T) {
	p1, p2, p3 := "fadkladfdsa", "fa12112fdsa", "A123fda!2*"
	assert.False(t, isValidPassword(&p1))
	assert.False(t, isValidPassword(&p2))
	assert.True(t, isValidPassword(&p3))
}

func TestRandomPassword(t *testing.T) {
	testTimes := 10000
	for i := 0; i < testTimes; i++ {
		password := RandomPassword(20)
		t.Logf("The random password is: %s", *password)
		assert.True(t, isValidPassword(password))
	}
}
