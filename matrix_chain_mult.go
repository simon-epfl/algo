package main

func matrixChainMult(dimensions []int) int {

	n := len(dimensions) - 1
	minimum := make([][]int, n+1)

	for i := 0; i <= n; i++ {
		minimum[i] = make([]int, n+1) // stockera le coût minimal (nombre de multiplications scalaires) pour calculer le produit Ai ... Aj

		minimum[i][i] = 0 // La multiplication d'une matrice avec elle-même coûte 0
	}

	for l := 2; l <= n; l++ { // pour chaque longueur de chaîne de matrices (on construit les petites chaînes en premier)
		for i := 1; i <= n-l+1; i++ { // i est l'indice de la première matrice de la chaîne
			j := i + l - 1 // j est l'indice de la dernière matrice de la chaîne
			minimum[i][j] = 9999999
			for k := i; k <= j-1; k++ { // toutes les paires (i, j) de matrices
				// On calcule le coût de la multiplication des matrices Ai...Ak et Ak+1...Aj
				q := minimum[i][k] + minimum[k+1][j] + dimensions[i-1]*dimensions[k]*dimensions[j]
				// p[i-1] est la dimension de la matrice Ai, p[k] est la dimension de la matrice Ak, et p[j] est la dimension de la matrice Aj
				if q < minimum[i][j] {
					// Si le coût est inférieur à celui précédemment enregistré, on le met à jour
					minimum[i][j] = q
				}
			}
		}
	}

	// m[1][n] contient le coût minimal pour multiplier les matrices A1...An
	return minimum[1][n]
}

func runMatrixChain() {
	dimensions := []int{30, 35, 15, 5, 10, 20, 25}
	result := matrixChainMult(dimensions)
	println("Minimum number of multiplications is:", result)
}
