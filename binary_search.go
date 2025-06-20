package main

func binarySearchAux(x int, array []int, begin int, end int) bool {
	// end est exclusif

	if begin >= end {
		return false // forcément vide
	}

	if end-begin == 1 {
		return x == array[begin] // un seul élément
	}

	middle := (begin + end) / 2
	if x <= array[middle] {
		return binarySearchAux(x, array, begin, middle+1)
	}
	return binarySearchAux(x, array, middle+1, end)
}

func runBinarySearch() {

	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	x := 10
	found := binarySearchAux(x, array, 0, len(array))
	if found {
		println("Element found:", x)
	} else {
		println("Element not found:", x)
	}

}
