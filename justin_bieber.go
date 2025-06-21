package main

/*
Restaurant placement. Justin Bieber has surprisingly decided to open a series of
restaurants along the highway between Geneva and Bern. The n possible locations are along a
straight line, and the distances of these locations from the start of the highway in Geneva are,
in kilometers and in arbitrary order, m1, m2, . . . , mn. The constraints are as follows:
• At each location, Justin may open at most one restaurant. The expected profit from
opening a restaurant at location i is pi, where pi > 0 and i = 1, 2, . . . , n.
• Any two restaurants should be at least k kilometers apart, where k is a positive integer.
As Justin is not famous for his algorithmic skills, he needs your help to find an optimal solution,
i.e., design and analyze an efficient algorithm to compute the maximum expected total profit
subject to the given constraints.
*/

func maxProfit(locations []int, profits []int, k int) int {

	// on considère que locations/profits sont déjà triés
	// mergesort(locations, 0, len(locations))

	cache := make([]int, len(locations))
	maxProfitSoFar := 0

	for i := 0; i < len(locations); i++ {
		cache[i] = profits[i]
		for j := 0; j < i; j++ {
			if locations[i]-locations[j] >= k {
				cache[i] = max(cache[i], profits[i]+cache[j])
			}
		}
		maxProfitSoFar = max(maxProfitSoFar, cache[i])
	}

	return maxProfitSoFar

}

func maxProfitFast(locations []int, profits []int, k int) int {

	nextAvailableLocation := make([]int, len(locations))

	// correction très smart : on garde un unique pointeur qu'on déplace
	lastNextAvailableLocation := 0
	for i := 0; i < len(locations); i++ {
		for lastNextAvailableLocation != len(locations)-1 &&
			locations[lastNextAvailableLocation] < locations[i]+k {
			lastNextAvailableLocation++
		}

		nextAvailableLocation[i] = lastNextAvailableLocation
	}

	cache := make([]int, len(locations)+1)
	cache[len(locations)] = 0

	for i := len(locations) - 1; i >= 0; i-- {
		cache[i] = max(cache[i+1], profits[i]+cache[nextAvailableLocation[i]])
	}

	return cache[0]

}

func runMaxProfit() {
	locations := []int{1, 3, 5, 10, 15}
	profits := []int{5, 6, 10, 3, 12}
	k := 4

	maxProfitValue := maxProfitFast(locations, profits, k)
	println("Maximum profit:", maxProfitValue)
}
