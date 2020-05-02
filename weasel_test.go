package weasel

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitial(test *testing.T) {
	rand.Seed(1)

	gotText := Initial(5)

	assert.Equal(test, "OPCLE", gotText)
}

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

func TestSearch_withoutSample(test *testing.T) {
	variants := []string{"o1e t2p 3hree", "o21 2wo t3ree", "15e 133 t34e1"}
	sample := "one two three"
	result := Search(variants, sample)

	assert.Equal(test, "o1e t2p 3hree", result)
}

func TestSearch_withSample(test *testing.T) {
	variants := []string{"one two three", "o21 2wo t3ree", "o5e 133 t34e1"}
	sample := "one two three"
	result := Search(variants, sample)

	assert.Equal(test, sample, result)
}
