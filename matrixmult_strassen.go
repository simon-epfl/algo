package main

import "fmt"

/*
> [!tip] Sous-problème plus simple, multiplier des nombres complexes
>
> $(a + i b) dot (c + i d) = (a dot c - b dot d) + i (a dot d + b dot c) = r$
>
> Pour multiplier 2 complexes, on a dû faire 4 produits de nombres réels.
> Calculons :
> $s_1 = (a + b) dot (c + d) = a c + a d + b c + b d$
> $s_2 = a dot c$
> $s_3 = b dot d$
>
> $r = (s_2 - s_3) + i(s_1 - s_2 - s_3)$
>
> Maintenant on a un produit de moins nécessaire pour trouver $r$! On a plus d'additions et de soustractions. On peut utiliser le même principe pour les matrices.

/*
On a comme entrées deux matrices carrées, $n times n$ :
- $A = (a_(i j))$
- $B = (b_(i j))$

On sort une matrice carrée $n times n$ : $C = (c_(i j))$ où $A dot B = C$.

Example (n = 2)

$$ mat(c_(1 1), c_(1 2); c_(2 1), c_(2 2) ) = mat(a_11, b_11; a_21, a_22) dot mat(b_11, b_12; b_21, b_22) $$
$$c_11 = a_11 b_11 + a_12 b_21 + ... + a_(1 n) b_(n 1) = sum_(k = 1)^n a_(1 k) b_(k 1)$$
$$ " Plus généralement, " c_(i j) = sum_(k = 1)^n a_(i k)b_(k j) $$

On peut écrire un algo simple qui en temps $Theta(n^3)$ qui calcule $c_(i j)$ (trois boucles for jusqu'à $n$ qui pour chaque entrée $i, j$ somme tous les produits).
$arrow$ le temps utilisé par cet algo est $Theta(n^2)$, parce qu'on ne créé par une variable à chaque boucle (pour la dernière boucle qui somme les produits on les ajoute à une variable existante $c_(i j)$).
*/

func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		for _, val := range row {
			fmt.Print(val, " ")
		}
		fmt.Println()
	}
}

func squareMatricesNotEmptySameSize(matrixA [][]int, matrixB [][]int) bool {
	if len(matrixA) == 0 || len(matrixB) == 0 || len(matrixA) != len(matrixB) {
		return false
	}
	return true
}

func addmatrix(matrixA [][]int, matrixB [][]int) [][]int {
	if !squareMatricesNotEmptySameSize(matrixA, matrixB) {
		return nil
	}

	result := make([][]int, len(matrixA))
	for i := range matrixA {
		result[i] = make([]int, len(matrixA[0]))
		for j := range matrixA[0] {
			result[i][j] = matrixA[i][j] + matrixB[i][j]
		}
	}
	return result
}

func splitmatrix4(matrix [][]int) ([][]int, [][]int, [][]int, [][]int) {
	n := len(matrix)

	a11 := make([][]int, n/2)
	a12 := make([][]int, n/2)
	a21 := make([][]int, n/2)
	a22 := make([][]int, n/2)

	for i := 0; i < n/2; i++ { // pour chaque ligne de la matrice d'origine
		a11[i] = make([]int, n/2) // comment on veut une matrice carrée, on initialise les lignes avec n/2 colonnes
		a12[i] = make([]int, n/2)
		a21[i] = make([]int, n/2)
		a22[i] = make([]int, n/2)
		for j := 0; j < n/2; j++ { // et pour chaque colonne de la matrice d'origine, on remplit la ligne
			a11[i][j] = matrix[i][j]
			a12[i][j] = matrix[i][j+n/2]
			a21[i][j] = matrix[i+n/2][j]
			a22[i][j] = matrix[i+n/2][j+n/2]
		}
	}

	return a11, a12, a21, a22
}

func matrixmult(matrixA [][]int, matrixB [][]int) [][]int {
	if !squareMatricesNotEmptySameSize(matrixA, matrixB) {
		return nil
	}

	// on a atteint une matrice de taille 1x1! on ne multiplie plus les matrices, on retourne le produit
	if (len(matrixA) == 1) && (len(matrixA[0]) == 1) {
		return [][]int{{matrixA[0][0] * matrixB[0][0]}}
	}

	a11, a12, a21, a22 := splitmatrix4(matrixA)
	b11, b12, b21, b22 := splitmatrix4(matrixB)

	M1 := matrixmult(addmatrix(a11, a22), addmatrix(b11, b22))
	M2 := matrixmult(addmatrix(a21, a22), b11)
	M3 := matrixmult(a11, addmatrix(b12, b22))
	M4 := matrixmult(a22, addmatrix(b21, b11))
	M5 := matrixmult(addmatrix(a11, a12), b22)
	M6 := matrixmult(addmatrix(a21, a11), b11)
	M7 := matrixmult(addmatrix(a12, a22), b21)

	c11 := addmatrix(addmatrix(M1, M4), addmatrix(M7, M5))
	c12 := addmatrix(M3, M5)
	c21 := addmatrix(M2, M4)
	c22 := addmatrix(addmatrix(M1, M3), addmatrix(M6, M2))

	// on reconstruit la matrice C à partir des sous-matrices c11, c12, c21, c22
	C := make([][]int, len(matrixA))
	for i := range C {
		C[i] = make([]int, len(matrixA[0]))
	}
	for i := 0; i < len(matrixA)/2; i++ {
		for j := 0; j < len(matrixA)/2; j++ {
			C[i][j] = c11[i][j]
			C[i][j+len(matrixA)/2] = c12[i][j]
			C[i+len(matrixA)/2][j] = c21[i][j]
			C[i+len(matrixA)/2][j+len(matrixA)/2] = c22[i][j]
		}
	}

	return C
}

func strassen(matrixA [][]int, matrixB [][]int) [][]int {
	if !squareMatricesNotEmptySameSize(matrixA, matrixB) {
		return nil
	}
	return matrixmult(matrixA, matrixB)
}

func runstrassen() {
	matrixA := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	matrixB := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	result := strassen(matrixA, matrixB)
	fmt.Println("Résultat de la multiplication de matrices Strassen:")
	printMatrix(result)
}
