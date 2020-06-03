package main

import (
	"time"

	"camel.uangel.com/ua5g/ulib.git/exec"
	"camel.uangel.com/ua5g/ulib.git/ulog"
)

func main() {
	err := doSomethingWithFuture()
	if err != nil {
		ulog.Fatal("Failed to do something: %s", err)
	}
}

func doSomethingWithFuture() error {

	ch := make(chan error, 1)

	execContext := exec.ContextNoLimit()
	queryTimeout, _ := time.ParseDuration("1s")

	future := exec.GoWithTimeout(execContext, queryTimeout, func() {
		err := doSomething()
		// doSomething이 완료되면 결과를 ch 로 보냅니다.
		// 타임 아웃이 먼저 발생해도 , 여기는 항상 실행된다는 것을 주의 하세요.
		// 그래서  ch 의 크기를 1로 잡거나, 여기서는 non blocking으로 send 해야 합니다.
		ch <- err

		// exec.Context 에 discard 기능이 있다면 , 아예 실행이 안될수도 있다는 점도 염두에 두세요.
	})

	future.OnFailure(func(err error) {
		ch <- err
	}, exec.ContextNoLimit())

	err := <-ch
	ulog.Info("error=%#v", err)

	return err
}

func doSomething() error {
	time.Sleep(2 * time.Second)
	return nil
}
