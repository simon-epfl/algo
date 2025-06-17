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

	c := make([][]int, len1+1) // il stockera c[index_de_fin_de_str1][index_de_fin_de_str2] = #size of longest subsequence
	for i := 0; i <= len1; i++ {
		c[i] = make([]int, len2+1)
		c[i][0] = 0
	}

	for j := 0; j <= len2; j++ {
		c[0][j] = 0
	}

	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {

			//fmt.Printf("testing %c %c", str1[i-1], str2[j-1])
			//fmt.Printf("i=%d j=%d", i, j)

			lcsShorten1 := c[i-1][j]
			lcsShorten2 := c[i][j-1]

			if str1[i-1] == str2[j-1] {
				c[i][j] = max(
					lcsShorten1,
					lcsShorten2,
					1+c[i-1][j-1],
				)
			} else {
				c[i][j] = max(
					lcsShorten1,
					lcsShorten2,
				)
			}

			//fmt.Println("found that current is", c[i][j])
		}
	}

	return c[len1][len2]
}

func runLongestCommonSubsequence() {
	str1 := "abcde"
	str2 := "ace"

	result := longestCommonSubsequenceFast(str1, str2)
	println("Length of LCS is:", result)
}
