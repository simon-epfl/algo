package main

func fastPower(current int, exp int) int {
	if exp == 0 {
		return 1
	}

	if exp%2 == 0 {
		half := fastPower(current, exp/2)
		return half * half
	} else {
		return current * fastPower(current, exp-1)
	}
}

func runFastPower() {
	base := 2
	exp := 23

	result := fastPower(base, exp)
	println("Result of fast power:", result)
}
