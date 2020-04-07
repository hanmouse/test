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

	numCPU = runtime.NumCPU()
	fmt.Println("numCPU:", numCPU)

	// 0: query
	numCPU = runtime.GOMAXPROCS(0)
	fmt.Println("numCPU:", numCPU)
}
