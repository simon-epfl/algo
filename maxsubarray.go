package main

import "math"

func maxsubarray(array []int, begin int, end int) int {

	if begin < end {
		middle := int(math.Floor(float64(begin+end) / 2))

		maxCrossingLeft := 0
		currCrossingLeft := 0
		for i := middle; i >= begin; i-- {
			currCrossingLeft += array[i]
			if currCrossingLeft > maxCrossingLeft {
				maxCrossingLeft = currCrossingLeft
			}
		}

		maxCrossingRight := 0
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

func runmaxsubarray() {

	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	result := maxsubarray(nums, 0, len(nums)-1)
	println("La somme maximale d'un sous-tableau est:", result)
}
