package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	weasel "github.com/irenicaa/go-weasel"
)

const outputRate = 10

func main() {
	rand.Seed(time.Now().UnixNano())

	sample := flag.String("sample", "METHINKS IT IS LIKE A WEASEL", "")
	rate := flag.Float64("rate", 0.05, "")
	population := flag.Int("population", 100, "")
	flag.Parse()

	generation := 0
	current := weasel.Initial(len(*sample))
	for ; weasel.Fitness(current, *sample) > 0; generation++ {
		if generation%outputRate == 0 {
			fmt.Println(generation, current)
		}

		variants := weasel.Populate(current, *rate, *population)
		current = weasel.Search(variants, *sample)
	}

	fmt.Println(generation, current)
}
