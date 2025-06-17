package main

import "fmt"

func insertionsort(array []int) {

	for j := 1; j < len(array); j++ {
		key := array[j]

		fmt.Println("On insère l'élément", key, "à sa place dans le tableau")
		// on commence à comparer juste à gauche de notre clef (parce qu'on sait que le tableau est bien trié)
		i := j - 1
		for i >= 0 && array[i] > key {
			array[i+1] = array[i] // move the key to the right to make space for j
			i = i - 1
		}
		array[i+1] = key

		fmt.Println("Tableau après l'insertion de l'élément", key, ":", array)
	}

}

func runInsertionsort() {

	nums := []int{3, 6, 8, 10, 0, 2, 1}

	fmt.Println("Tableau avant le tri:", nums)

	insertionsort(nums)

	fmt.Println("Tableau trié:", nums)
}
