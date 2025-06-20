package main

import (
	"os"
	"time"
)

func main() {

	algoName := os.Args[1]
	startTime := time.Now()

	switch algoName {
	case "fast_power":
		runFastPower()
	case "kruskal":
		runKruskal()
	case "prims":
		runPrims()
	case "connected_components":
		connectedComponents()
	case "ford_fulkerson":
		fordFulkerson()
	case "strongly_connected_components":
		stronglyConnectedComponents()
	case "topological_sort":
		topologicalSort()
	case "dfs":
		runDfs()
	case "bfs":
		bfs()
	case "optimal_bst":
		runOptimalBst()
	case "changemaking":
		runChangemaking()
	case "fibotop":
		runFiboTopdown()
	case "fibobottom":
		runFiboBottomup()
	case "fibo":
		runDumbfibo()
	case "longest_common_seq":
		runLongestCommonSubsequence()
	case "matrix_chain_mult":
		runMatrixChain()
	case "rodcutting":
		runRodcutting()
	case "heapsort":
		runHeapsort()
	case "strassen":
		fallthrough
	case "matrixmult":
		fallthrough
	case "matrixmult_strassen":
		runStrassen()
	case "maxsubarray":
		runMaxsubarray()
	case "mergesort":
		runMergesort()
	case "insertionsort":
		runInsertionsort()
	case "linearsearch":
		runLinearsearch()
	case "dijkstra":
		dijkstra()
	case "bellmanford":
		bellmanFord()
	case "quicksort":
		runQuicksort()
	case "quickselect":
		quickselect()
	case "countingsort":
		countingsort()
	default:
		println("Unknown algorithm")
	}

	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	println("Exec:", elapsedTime.Microseconds(), "µs")

}
