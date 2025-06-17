package main

import (
	"os"
	"time"
)

func main() {

	algoName := os.Args[1]
	startTime := time.Now()

	switch algoName {
	case "optimal_bst":
		runOptimalBst()
	case "changemaking":
		runchangemaking()
	case "fibotop":
		runfibotopdown()
	case "fibobottom":
		runfibobottomup()
	case "fibo":
		rundumbfibo()
	case "longest_common_seq":
		runLongestCommonSubsequence()
	case "matrix_chain_mult":
		runMatrixChain()
	case "rodcutting":
		runrodcutting()
	case "heapsort":
		runheapsort()
	case "strassen":
		fallthrough
	case "matrixmult":
		fallthrough
	case "matrixmult_strassen":
		runstrassen()
	case "maxsubarray":
		runmaxsubarray()
	case "mergesort":
		runmergesort()
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

	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	println("Exec:", elapsedTime.Microseconds(), "µs")

}
