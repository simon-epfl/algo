package main

import "fmt"

func countingsort() {

	// Exemple d'utilisation
	nums := []int{2, 5, 3, 0, 2, 3, 0, 3}
	k := 5 // on a au plus k éléments différents

	count := make([]int, k+1) // C[0, .., k]

	result := make([]int, len(nums))

	// Initialiser le tableau de comptage
	for i := 0; i < k+1; i++ {
		count[i] = 0
	}

	// Compter les occurrences de chaque élément
	for i := 0; i < len(nums); i++ {
		count[nums[i]]++
	}

	fmt.Println("Tableau de comptage:", count)

	for i := 1; i < k+1; i++ {
		count[i] += count[i-1]
	}

	fmt.Println("Tableau de comptage cumulé:", count)

	// Construire le tableau trié
	for j := len(nums) - 1; j >= 0; j-- {
		result[count[nums[j]]-1] = nums[j]
		count[nums[j]]--
	}

	fmt.Println("Tableau trié:", result)

}
