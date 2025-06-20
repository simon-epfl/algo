package main

import (
	"math"
)

/*
On a une entrée $[0, 2, -4, 3, -1, 4, 5, 7, -9]$ et on veut trouver le sous-intervalle avec la plus grande somme $[3, -1, 4, 5, 7]$.

On peut tout bruteforce, tester toutes les combinaisons : $O(n^2)$.

On peut faire mieux en $n log n$ avec du divide et conquer :
- on peut séparer le problème en deux (c'est facile on coupe au milieu)
- mais comment combiner les deux, une fois qu'on a les solutions de l'un et de l'autre ? (en reprenant l'exemple plus haut, $[3], [5,7]$).
	- on connaît la solution à gauche
	- on connaît la solution à droite
	- mais on doit vérifier qu'en se chevauchant on a pas une meilleure solution $O(n)$:
		- pour ça on doit trouver la plus grande somme à gauche qui commence au milieu (ici $3 - 1 = 2$)
		- trouver la plus grande somme à droite qui commence au milieu (ici $4 + 5 + 7 = 16$)
		- les sommer ($2 + 16 = 18$)
	- comparer ces trois solutions prendre la meilleure
*/

func maxsubarray(array []int, begin int, end int) int {

	if begin < end {
		middle := int(math.Floor(float64(begin+end) / 2))

		maxCrossingLeft := int(math.MinInt64)
		currCrossingLeft := 0
		for i := middle; i >= begin; i-- {
			currCrossingLeft += array[i]
			if currCrossingLeft > maxCrossingLeft {
				maxCrossingLeft = currCrossingLeft
			}
		}

		maxCrossingRight := int(math.MinInt64)
		currCrossingRight := 0
		for i := middle + 1; i <= end; i++ {
			currCrossingRight += array[i]
			if currCrossingRight > maxCrossingRight {
				maxCrossingRight = currCrossingRight
			}
		}

		return max(
			maxsubarray(array, begin, middle),
			maxCrossingLeft+maxCrossingRight,
			maxsubarray(array, middle+1, end),
		)

	}

	return array[begin]
}

func maxsubarray_linear(array []int) int {

	maxSoFar := -1
	maxEndingHere := -1

	for _, item := range array {

		maxEndingHere = max(maxEndingHere+item, item)
		maxSoFar = max(maxEndingHere, maxSoFar)

	}

	return maxSoFar
}

func runMaxsubarray() {

	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	result := maxsubarray(nums, 0, len(nums)-1)
	println("La somme maximale d'un sous-tableau est:", result)
}
