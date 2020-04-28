package main

import (
	"fmt"
	"time"
)

func main() {

	var sleepTime time.Duration = 3 * time.Second

	startTime := time.Now()
	fmt.Printf("Timer started: %v (waiting for %v)\n", startTime.Format(time.UnixDate), sleepTime)

	timeAfter := <-time.After(sleepTime)

	fmt.Printf("Timer stopped: %v (%v elapsed)\n", timeAfter.Format(time.UnixDate), time.Since(startTime))
}
