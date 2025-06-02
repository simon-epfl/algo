package main

import (
	"os"
)

func main() {

	algoName := os.Args[1]

	switch algoName {
	case "insertionsort":
		runinsertionsort()
	case "linearsearch":
		runlinearsearch()
	case "dijkstra":
		dijkstra()
	case "bellmanford":
		bellmanford()
	case "quicksort":
		runquicksort()
	case "quickselect":
		quickselect()
	case "countingsort":
		countingsort()
	default:
		println("Unknown algorithm")
	}

}
