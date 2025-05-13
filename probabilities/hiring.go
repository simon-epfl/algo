package main

import (
	"fmt"
	"math"
	"math/rand"
)

func hiringProblem(n int) int {
	// On génère un nombre de candidats (0 = best, n-1 = worst)
	candidates := rand.Perm(n)

	// Determine sample size (n/e)
	sampleSize := int(float64(n) / math.E)
	bestSoFar := n + 1 // initialize to something worse than worst

	// Step 3: Observe first `sampleSize` candidates
	for i := 0; i < sampleSize; i++ {
		if candidates[i] < bestSoFar {
			bestSoFar = candidates[i]
		}
	}

	// Step 4: Choose the first candidate better than bestSoFar
	for i := sampleSize; i < n; i++ {
		if candidates[i] < bestSoFar {
			return candidates[i] // return rank of hired candidate
		}
	}

	return candidates[n-1] // if none are better, hire the last
}

func main() {

	n := 100
	simulations := 100000
	success := 0

	for i := 0; i < simulations; i++ {
		rank := hiringProblem(n)
		if rank == 0 { // we hired the best
			success++
		}
	}

	fmt.Printf("Success rate of hiring the best: %.2f%%\n", float64(success)/float64(simulations)*100)
}
