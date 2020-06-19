package main

import (
	"fmt"
	"time"
)

func main() {
	//afterTest()
	//timerTest()
	afterFuncTest()
}

func afterTest() {
	var sleepTime time.Duration = 3 * time.Second

	startTime := time.Now()
	fmt.Printf("Timer started: %v (waiting for %v)\n", startTime.Format(time.UnixDate), sleepTime)

	timeAfter := <-time.After(sleepTime)

	fmt.Printf("Timer stopped: %v (%v elapsed)\n", timeAfter.Format(time.UnixDate), time.Since(startTime))
}

func timerTest() {

	timer1 := time.NewTimer(time.Second * 2)
	<-timer1.C
	fmt.Println("Timer 1 expired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()

	timer2Stopped := timer2.Stop()
	if timer2Stopped {
		fmt.Println("Timer 2 stopped")
	}

	timer2Stopped = timer2.Stop()
	if !timer2Stopped {
		fmt.Println("Timer 2 already stopped")
	}
}

func afterFuncTest() {

	timeout := 1 * time.Second

	time.AfterFunc(timeout,
		// 이 함수는 알아서 go routine으로 실행된다.
		func() {
			fmt.Printf("I'm go routine. Timer expired (timeout: %s)\n", timeout.String())
		},
	)

	fmt.Printf("I'm main. Timer started (timeout: %s)\n", timeout.String())

	// time.AfterFunc()의 f가 실행될 때까지 대기한다.
	time.Sleep(timeout + time.Second)
}
