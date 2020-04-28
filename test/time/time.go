package main

import (
	"fmt"
	"time"
)

func main() {

	var sleepTime time.Duration = 3 * time.Second

	fmt.Printf("Timer started (waiting for %v): %v\n", sleepTime, time.Now().Format(time.UnixDate))

	timeAfter := <-time.After(sleepTime)

	fmt.Printf("Timer stopped: %v\n", timeAfter.Format(time.UnixDate))
}
