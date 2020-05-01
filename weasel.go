package weasel

import "math/rand"

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ "

// Mutate ...
func Mutate(text string, rate float64) string {
	textCopy := ""
	for i := 0; i < len(text); i++ {
		if rand.Float64() > rate {
			textCopy += string(text[i])
		} else {
			textCopy += string(alphabet[rand.Intn(len(alphabet))])
		}
	}

	return textCopy
}

// Populate ...
func Populate(text string, rate float64, count int) []string {
	textGroup := []string{}
	for i := 0; i < count; i++ {
		textGroup = append(textGroup, Mutate(text, rate))
	}

	return textGroup
}
