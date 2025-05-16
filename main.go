package main

import "os"

func main() {

	algoName := os.Args[1]

	switch algoName {
	case "dijkstra":
		dijkstra()
	case "bellmanford":
		bellmanford()
	case "quicksort":
		runquicksort()
	default:
		println("Unknown algorithm")
	}

}
