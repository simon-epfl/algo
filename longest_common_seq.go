package main

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
			longestCommonSubsequence(str1[0:len1-1], str2),
			longestCommonSubsequence(str1, str2[0:len2-1]),
		)
	}
}

func longestCommonSubsequenceFast(str1 string, str2 string) int {

	len1 := len(str1)
	len2 := len(str2)

	c := make([][]int, len1) // il stockera c[index_de_fin_de_str1][index_de_fin_de_str2] = longest subsequence
	for i := 0; i < len1; i++ {
		c[i] = make([]int, len2)
		c[i][0] = 0
	}

	for j := 0; j < len2; j++ {
		c[0][j] = 0
	}

	// todo
	return -1
}

func runLongestCommonSubsequence() {
	str1 := "hsod"
	str2 := "dwjh"

	result := longestCommonSubsequence(str1, str2)
	println("Length of LCS is:", result)
}
