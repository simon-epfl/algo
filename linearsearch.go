package main

func linearsearch(array []int, key int) int {
	for i := 0; i < len(array); i++ {
		if array[i] == key {
			return i
		}
	}
	return -1
}

func runLinearsearch() {

	array := []int{3, 6, 8, 10, 0, 2, 1}
	key := 10
	result := linearsearch(array, key)
	if result != -1 {
		println("L'élément", key, "a été trouvé à l'index", result)
	} else {
		println("L'élément", key, "n'a pas été trouvé dans le tableau")
	}
}
