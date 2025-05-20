package main

import "fmt"

func quickselectHelper(array []int, low, high, k int) int {
	if low == high {
		return array[low]
	}

	pivotIndex := betterPartitionRandomPivot(array, low, high)

	// on a exactement k éléments à gauche du pivot!
	if k == pivotIndex {
		return array[k]
	} else if k < pivotIndex {
		// on cherche le kème plus petit élément dans la partie gauche
		return quickselectHelper(array, low, pivotIndex-1, k)
	} else {
		// on cherche le kème plus petit élément dans la partie droite
		return quickselectHelper(array, pivotIndex+1, high, k)
	}

}

func quickselect() {

	nums := []int{3, 6, 8, 10, 0, 2, 1}

	k := 4 // 6ème plus petit élément

	result := quickselectHelper(nums, 0, len(nums)-1, k-1)
	fmt.Println("Le", k, "ème plus petit élément est:", result)
}
