package main

import "fmt"

// very slow, no cache
func longestCommonSubsequence(str1 string, str2 string) int {

	len1 := len(str1)
	len2 := len(str2)

	if len1 == 0 || len2 == 0 {
		return 0
	}

	lastChar1 := str1[len1-1]
	lastChar2 := str2[len2-1]

	if lastChar1 == lastChar2 {
		return max(
			longestCommonSubsequence(str1[0:len1-1], str2),
			longestCommonSubsequence(str1, str2[0:len2-1]),
			1+longestCommonSubsequence(str1[0:len1-1], str2[0:len2-1]),
		)
	} else {
		return max(
			longestCommonSubsequence(str1[0:len1-1], str2), // D BAB 0
			longestCommonSubsequence(str1, str2[0:len2-1]), // DA BA 1
		)
	}
}

type Direction int

const (
	Left Direction = iota + 1
	Up
	Diag
)

func longestCommonSubsequenceFast(str1 string, str2 string) (int, [][]Direction) {

	len1 := len(str1)
	len2 := len(str2)

	cache := make([][]int, len1+1) // il stockera c[index_de_fin_de_str1][index_de_fin_de_str2] = #size of longest subsequence
	direction := make([][]Direction, max(len1, len2)+1)
	for i := 0; i <= len1; i++ {
		cache[i] = make([]int, len2+1)
		cache[i][0] = 0

		direction[i] = make([]Direction, len2+1)
	}

	for j := 0; j <= len2; j++ {
		cache[0][j] = 0
	}

	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {

			//fmt.Printf("testing %c %c", str1[i-1], str2[j-1])
			//fmt.Printf("i=%d j=%d", i, j)

			lcsShorten1 := cache[i-1][j]
			lcsShorten2 := cache[i][j-1]

			if str1[i-1] == str2[j-1] {
				cache[i][j] = max(
					lcsShorten1,
					lcsShorten2,
					1+cache[i-1][j-1],
				)
			} else {
				cache[i][j] = max(
					lcsShorten1,
					lcsShorten2,
				)
			}

			switch cache[i][j] {
			case lcsShorten1:
				direction[i][j] = Left
			case lcsShorten2:
				direction[i][j] = Up
			default:
				direction[i][j] = Diag
			}

			//fmt.Println("found that current is", c[i][j])
		}
	}

	return cache[len1][len2], direction
}

func runLongestCommonSubsequence() {
	str1 := "dac"
	str2 := "bab"

	result, directions := longestCommonSubsequenceFast(str1, str2)
	println("Length of LCS is:", result)

	for i := 0; i <= len(str1); i++ {
		for j := 0; j <= len(str2); j++ {
			fmt.Printf("%d ", directions[i][j])
		}
		fmt.Println()
	}

	str := ""
	lastI := len(str1)
	lastJ := len(str2)
	for lastI > 0 && lastJ > 0 {
		dir := directions[lastI][lastJ]
		//fmt.Printf("lastI=%d lastJ=%d dir=%d\n", lastI, lastJ, dir)
		switch dir {
		case Diag:
			{
				str += string(str1[lastI-1])
				//fmt.Println("adding", string(str1[lastI-1]), "to str")
				lastI--
				lastJ--
			}
		case Left:
			{
				lastI--
			}
		case Up:
			{
				lastJ--
			}
		}
	}

	// reverse the string
	for i := len(str) - 1; i >= 0; i-- {
		fmt.Print(string(str[i]))
	}
	fmt.Println()
}
