package main

import (
	"fmt"
	"math/rand"
)

// on veut partitionner le tableau en deux parties
// cette fonction réorganise le tableau de telle sorte à ce que les éléments
// inférieurs ou égaux au pivot soient à gauche et les éléments supérieurs soient à droite
func partition(array []int, low, high int) int {

	// on choisit le dernier élément comme pivot (arbitrairement, plus tard ce sera random!)
	pivot := array[high]

	// i représentera le dernier élément de la zone contenant des éléments ≤ pivot
	i := low - 1

	// on parcourt le tableau et on compare chaque élément avec le pivot
	for j := low; j < high; j++ {
		// l'élément courant est inférieur ou égal au pivot!
		if array[j] <= pivot {
			// on agrandit la zone contenant des éléments ≤ pivot
			// donc i pointe maintenant vers l'élément juste en dehors de la zone
			i++
			// on échange l'élément courant avec cet élément juste en dehors de la zone
			// donc i pointe maintenenant correctement vers le dernier élément de la zone! (array[j])
			array[i], array[j] = array[j], array[i]
		}
	}

	array[i+1], array[high] = array[high], array[i+1] // on insert le pivot à sa place finale
	return i + 1                                      // la position du pivot
}

// problème : le nombre d'appels récursifs dépend du tableau et du pivot!
// si le tableau est déjà trié, on va faire n appels récursifs!
// solution : choisir un pivot aléatoire
func betterPartitionRandomPivot(array []int, low, high int) int {
	// on choisit un pivot aléatoire entre low et high
	pivotIndex := rand.Intn(high-low+1) + low
	// on échange le pivot avec le dernier élément
	array[pivotIndex], array[high] = array[high], array[pivotIndex]
	return partition(array, low, high) // on appelle la fonction de partitionnement
}

// slide 8
// https://moodle.epfl.ch/pluginfile.php/3442098/mod_resource/content/1/Lecture24.pdf

func quicksortRecursive(array []int, low, high int) {
	if low >= high {
		return
	}
	pivotIndex := betterPartitionRandomPivot(array, low, high)
	quicksortRecursive(array, low, pivotIndex-1)
	quicksortRecursive(array, pivotIndex+1, high)
}

func quicksort(array []int) {
	quicksortRecursive(array, 0, len(array)-1)
}

func runQuicksort() {
	array := []int{5, 8, 4, 7, 1, 2, 3, 6}
	fmt.Println("Original array:", array)
	quicksort(array)
	fmt.Println("Sorted array:", array)
}
