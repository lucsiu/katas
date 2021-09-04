package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSwitch(t *testing.T) {
	assert.Equal(t, 0, GetPlayerNumber("Luca"))
	assert.Equal(t, 1, GetPlayerNumberV2("allen"))

	assert.Equal(t, "I am a string.\n", V3("allen"))
	assert.Equal(t, "I am a int.\n", V3(1))
	assert.Equal(t, "I am a bool.\n", V3(false))
	assert.Equal(t, "Do not know typo float64.\n", V3(12.21))
}
