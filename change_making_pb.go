package main

// pour tous les types de coins i
// 1 + c(n - i)

func changemaking(amount int, prices []int, cache []int) int {

	if cache[amount] > -9999 {
		return cache[amount]
	}

	current := 99999
	for i := 0; i < len(prices); i++ {
		if (amount - prices[i]) < 0 {
			continue // on ne peut pas utiliser cette pièce
		}
		current = min(
			1+changemaking(amount-prices[i], prices, cache),
			current,
		)
	}

	cache[amount] = current

	return cache[amount]

}

func runChangemaking() {

	amount := 11
	prices := []int{2, 5}

	cache := make([]int, amount+1)
	for i := range cache {
		cache[i] = -9999
	}
	cache[0] = 0 // pour un montant de 0, on n'a pas besoin de pièces

	result := changemaking(amount, prices, cache)

	if result >= 99999 {
		println("Impossible de faire la monnaie pour le montant:", amount)
	} else {
		println("Le nombre minimum de pièces pour faire la monnaie de", amount, "est:", result)
	}

}
