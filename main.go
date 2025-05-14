package main

import "fmt"

func main() {
	array := []int{5, 8, 4, 7, 1, 2, 3, 6}
	fmt.Println("Array before sorting:", array)
	quicksort(array)
	fmt.Println("Array after sorting:", array)

}
