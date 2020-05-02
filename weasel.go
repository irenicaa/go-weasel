package weasel

import "math/rand"

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ "

// Initial ..
func Initial(length int) string {
	text := ""
	for i := 0; i < length; i++ {
		text += string(alphabet[rand.Intn(len(alphabet))])
	}

	return text
}

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

// Fitness ...
func Fitness(text string, sample string) int {
	count := 0
	for i := 0; i < len(text); i++ {
		if text[i] != sample[i] {
			count++
		}
	}

	return count
}

// Search ...
func Search(variants []string, sample string) string {
	minCount := len(sample)
	result := ""
	for i := 0; i < len(variants); i++ {
		count := Fitness(variants[i], sample)
		if count < minCount {
			minCount = count
			result = variants[i]
		}
	}

	return result
}
