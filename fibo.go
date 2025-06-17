package main

/*
➜  algo git:(main) ✗ go build -o algo && ./algo fibo
Fibonacci de 40 : 165580141
Exec: 468230 µs
➜  algo git:(main) ✗ go build -o algo && ./algo fibotop
Fibonacci de 40 : 165580141
Exec: 10 µs
➜  algo git:(main) ✗ go build -o algo && ./algo fibodown
Unknown algorithm
Exec: 6 µs
*/

func dumbFibo(n int) int {

	if n == 0 || n == 1 {
		return 1
	}

	return dumbFibo(n-1) + dumbFibo(n-2)
}

func fiboBottomUp(n int) int {

	cache := make([]int, n)

	cache[0] = 1
	cache[1] = 1

	if n < 2 {
		return cache[n]
	}

	for i := 2; i <= n; i++ {
		cache[i] = cache[i-1] + cache[i-2]
	}

	return cache[n]

}

func fiboTopDown(n int, cache []int) int {

	if cache[n] > -9999 {
		return cache[n]
	}

	if n < 2 {
		cache[0] = 1
		cache[1] = 1
		return 1
	}

	res := fiboTopDown(n-1, cache) + fiboTopDown(n-2, cache)

	cache[n] = res

	return res

}

func runfibotopdown() {
	n := 40
	cache := make([]int, n+1)
	for i := range cache {
		cache[i] = -9999
	}

	result := fiboTopDown(n, cache)
	println("Fibonacci de", n, ":", result)
}

func runfibobottomup() {
	n := 40

	result := fiboBottomUp(n)
	println("Fibonacci de", n, ":", result)
}

func rundumbfibo() {
	n := 40

	result := dumbFibo(n)
	println("Fibonacci de", n, ":", result)
}
