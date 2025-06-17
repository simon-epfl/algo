package main

import "fmt"

func optimalBstAux(probabilities []int, cache [][]int, begin int, end int) int { // end is exclusive

	fmt.Println("computng optimalBstAux(", begin, ",", end, ")")

	if cache[begin][end] < 99999 {
		return cache[begin][end]
	}

	if begin == end { // on a un arbre vide
		cache[begin][end] = 0
		return 0
	}

	opti := 9999999

	for r := begin; r < end; r++ {
		sumProba := 0
		for i := begin; i < end; i++ {
			sumProba += probabilities[i]
		}
		leftTree := 0
		if r > begin { // on considère que notre arbre vaut zéro si on chosit comme root le premier élément
			leftTree = optimalBstAux(probabilities, cache, begin, r) // l'arbre gauche
		}
		rightTree := 0
		if r < end {
			rightTree = optimalBstAux(probabilities, cache, r+1, end) // l'arbre droit
		}
		current := leftTree + rightTree + sumProba // la somme des proba, parce qu'en fait on décale tout de +1 cran vers le bas
		// (dont le root lui-même, sa comparaison compte pour un accès)

		opti = min(current, opti)
	}

	fmt.Println("optimalBstAux(", begin, ",", end, ") = ", opti)

	cache[begin][end] = opti

	return opti

}

func optimalBst(probabilities []int) int {
	n := len(probabilities)
	cache := make([][]int, n+1)
	for i := range cache {
		cache[i] = make([]int, n+1)
		for j := range cache[i] {
			cache[i][j] = 99999 // on initialise le cache avec une valeur très haute
		}
	}

	return optimalBstAux(probabilities, cache, 0, n)
}

func runOptimalBst() {
	probabilities := []int{25, 20, 5, 20, 30} // les probabilités d'accès aux clés de l'arbre

	result := optimalBst(probabilities)

	println("Le coût minimum pour construire un arbre binaire de recherche optimal est:", (float64(result) / 100))
}
