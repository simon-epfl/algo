package main

func rodcuttingAux(prices []int, n int, cache []int) int {

	if n <= 0 {
		return 0
	}

	if cache[n-1] >= 0 {
		return cache[n-1]
	}

	q := -99999

	for i := 1; i < n; i++ {
		q = max(q, prices[i-1]+rodcuttingAux(prices, n-i, cache))
	}

	cache[n-1] = q

	return q

}

func rodcutting(prices []int, n int) int {

	cache := make([]int, n)

	for i := 0; i < n; i++ {
		cache[i] = -99999
	}

	cache[0] = 0 // prix pour une tige de longueur 0 est 0

	return rodcuttingAux(prices, n, cache)

}

func runrodcutting() {

	prices := []int{1, 5, 8, 9, 10, 17, 17, 20, 24, 30}

	n := len(prices)

	result := rodcutting(prices, n)

	println("Le prix maximum pour une tige de longueur", n, "est:", result)

}
