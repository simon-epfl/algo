package main

import (
	"fmt"
	"math"
)

func mergesort(array []int, begin int, end int) { // p est l'indice du début, r l'indice de fin
	if begin < end {
		middle := int(math.Floor(float64(begin+end) / 2))
		fmt.Println("On divise le tableau en deux parties begin=", begin, "middle=", middle, "end=", end)
		fmt.Print("Partie gauche:", array[begin:middle+1])
		fmt.Println("  | Partie droite:", array[middle+1:end+1]) // array[low:high], high est exclusif!
		mergesort(array, begin, middle)
		mergesort(array, middle+1, end)
		merge(array, begin, middle, end)
	}
}

func merge(array []int, begin int, middle int, end int) {
	n1 := middle - begin + 1 // la partie à gauche du tableau (triée)
	n2 := end - middle       // la partie à droite du tableau (triée)

	fmt.Println("merge(array, begin=", begin, ", middle=", middle, ", end=", end, ")")

	fmt.Println("On fusionne les parties gauche et droite du tableau")
	fmt.Println("Partie gauche:", array[begin:middle+1])
	fmt.Println("Partie droite:", array[middle+1:end+1])

	L := make([]int, n1+1)
	R := make([]int, n2+1)

	// on copie la partie à gauche dans le tableau
	for i := 0; i < n1; i++ {
		L[i] = array[begin+i]
	}

	// on copie la partie à droite dans le tableau
	for j := 0; j < n2; j++ {
		R[j] = array[middle+j+1]
	}

	fmt.Println("Tableau gauche après copie:", L[:n1])
	fmt.Println("Tableau droit après copie:", R[:n2])

	L[n1] = int(9999999)
	R[n2] = int(9999999)

	i := 0
	j := 0

	for k := begin; k <= end; k++ {
		if L[i] <= R[j] {
			// si L[i] est plus petit ou égal à R[j], on ajoute L[i] à l'index k
			// on passe à l'index suivant de L
			fmt.Println("On ajoute l'élément", L[i], "à l'index", k)
			array[k] = L[i]
			i = i + 1
		} else {
			// ça veut dire que cette fois-ci R[j] est plus petit que L[i], on ajoute R[j] à l'index k
			// on passe à l'index suivant de R
			fmt.Println("On ajoute l'élément", R[j], "à l'index", k)
			array[k] = R[j]
			j = j + 1
		}
	}

	fmt.Println("Tableau après fusion:", array)
}

func runmergesort() {
	nums := []int{3, 6, 8, 10, 0, 2, 1}

	fmt.Println("Tableau avant le tri:", nums)

	mergesort(nums, 0, len(nums)-1)

	fmt.Println("Tableau trié:", nums)
}
