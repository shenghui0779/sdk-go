package event

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignWithSHA1(t *testing.T) {
	sign := SignWithSHA1("2faf43d6343a802b6073aae5b3f2f109", "1606902086", "1246833592")

	assert.Equal(t, "ffb882ae55647757d3b807ff0e9b6098dfc2bc57", sign)
}
