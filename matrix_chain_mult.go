package main

func matrixChainMult(p []int) int {

	n := len(p) - 1
	m := make([][]int, n+1)
	s := make([][]int, n+1)

	for i := 0; i <= n; i++ {
		m[i] = make([]int, n+1) // stockera le coût minimal (nombre de multiplications scalaires) pour calculer le produit Ai ... Aj
		s[i] = make([]int, n+1) // stockera l'indice k qui minimise le coût de la multiplication

		m[i][i] = 0 // La multiplication d'une matrice avec elle-même coûte 0
	}

	for l := 2; l <= n; l++ { // l est la longueur de la chaîne de matrices
		for i := 1; i <= n-l+1; i++ {
			j := i + l - 1
			m[i][j] = 9999999
			for k := i; k <= j-1; k++ { // toutes les paires (i, j) de matrices
				// On calcule le coût de la multiplication des matrices Ai...Ak et Ak+1...Aj
				q := m[i][k] + m[k+1][j] + p[i-1]*p[k]*p[j]
				// p[i-1] est la dimension de la matrice Ai, p[k] est la dimension de la matrice Ak, et p[j] est la dimension de la matrice Aj
				if q < m[i][j] {
					// Si le coût est inférieur à celui précédemment enregistré, on le met à jour
					m[i][j] = q
					s[i][j] = k
				}
			}
		}
	}

	// m[1][n] contient le coût minimal pour multiplier les matrices A1...An
	return m[1][n]
}

func runMatrixChain() {
	p := []int{30, 35, 15, 5, 10, 20, 25}
	result := matrixChainMult(p)
	println("Minimum number of multiplications is:", result)
}
