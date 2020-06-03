package main

import (
	"time"

	"camel.uangel.com/ua5g/ulib.git/exec"
	"camel.uangel.com/ua5g/ulib.git/ulog"
)

func main() {

	var err error
	timeout, _ := time.ParseDuration("3s")
	runningTime, _ := time.ParseDuration("5s")

	err = doSomethingWithFuture(timeout, runningTime)
	if err != nil {
		ulog.Warn("Failed to do something with future: %s", err)
	}

	err = doSomethingWithAwaiter(timeout, runningTime)
	if err != nil {
		ulog.Warn("Failed to do something with awaiter: %s", err)
	}
}

func doSomethingWithFuture(timeout time.Duration, runningTime time.Duration) error {

	ch := make(chan error, 1)

	execContext := exec.ContextNoLimit()

	future := exec.GoWithTimeout(execContext, timeout, func() {
		err := doSomething(runningTime)
		// doSomething이 완료되면 결과를 ch 로 보냅니다.
		// 타임 아웃이 먼저 발생해도 , 여기는 항상 실행된다는 것을 주의 하세요.
		// 그래서  ch 의 크기를 1로 잡거나, 여기서는 non blocking으로 send 해야 합니다.
		ch <- err

		// exec.Context 에 discard 기능이 있다면 , 아예 실행이 안될수도 있다는 점도 염두에 두세요.
	})

	future.OnFailure(func(err error) {
		ch <- err
	}, exec.ContextNoLimit())

	return <-ch
}

func doSomething(runningTime time.Duration) error {
	time.Sleep(runningTime)
	return nil
}

func doSomethingWithAwaiter(timeout time.Duration, runningTime time.Duration) error {

	someContext := exec.ContextNoLimit()
	awaiter := exec.NewAwaiter(someContext, timeout)

	err := awaiter.AwaitFunc0Err(
		func() error {
			// 여기 안의 함수는 someContext 안에서 실행됩니다.
			err := doSomething(runningTime)
			return err
		},
	)

	return err
}
