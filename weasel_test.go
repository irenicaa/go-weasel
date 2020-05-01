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

func TestPopulate(test *testing.T) {
	rand.Seed(1)

	text := "the quick brown fox jumps over the lazy dog"
	gotText := Populate(text, 0.2, 3)

	wantText := []string{
		"the qu Pk brown fox jumps oveF tGD Nazy dog",
		"the N ick broRn foS jumps oveB theTlazyEdQg",
		"the quWck b Ewn foxMjumNs Qver She lRzy dog",
	}
	assert.Equal(test, wantText, gotText)
}

func TestFitness_withEqualStrings(test *testing.T) {
	text := "text text1 text2"
	sample := "text text1 text2"
	result := Fitness(text, sample)

	assert.Equal(test, 0, result)
}

func TestFitness_withNotEqualStrings(test *testing.T) {
	text := "text text2 text3"
	sample := "text text1 text2"
	result := Fitness(text, sample)

	assert.Equal(test, 2, result)
}
