package main

import (
	"fmt"
	"runtime"
)

func main() {
	cpuTest()
}

func cpuTest() {

	var numCPU int

	currMaxProcs := runtime.GOMAXPROCS(0)
	fmt.Println("currMaxProc:", currMaxProcs)

	numCPU = runtime.NumCPU()
	fmt.Println("numCPU:", numCPU)

	// 0: query
	runtime.GOMAXPROCS(numCPU - 1)
	currMaxProcs = runtime.GOMAXPROCS(0)
	fmt.Println("currMaxProc:", currMaxProcs)
}
