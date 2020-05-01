package weasel

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMutate(test *testing.T) {
	rand.Seed(1)

	text := "the quick brown fox jumps over the lazy dog"
	gotText := Mutate(text, 0.2)

	wantText := "the qu Pk brown fox jumps oveF tGD Nazy dog"
	assert.Equal(test, wantText, gotText)
}
