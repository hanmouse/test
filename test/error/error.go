package main

import (
	"errors"
	"fmt"
)

func main() {
	/*
		level := "MAJOR"
		err := fmt.Errorf("This is %v error", level)
		logger.Error(err.Error())

		fmt.Printf("err: %#v\n", err)
	*/

	// err는 a, b와 달리 여러 번 재할당 가능
	/*
		a, err := twoValuesReturningFunc1()
		fmt.Printf("a=%v, err=%s\n", a, err)

		b, err := twoValuesReturningFunc2()
		fmt.Printf("b=%v, err=%s\n", b, err)
	*/

	a, err := twoValuesReturningFunc1()
	if a == 1 {
		a, err := twoValuesReturningFunc1()
		fmt.Printf("a=%v, err=%s\n", a, err)
	}

	c, err := twoValuesReturningFunc1()
	fmt.Printf("c=%v, err=%s\n", c, err)
}

func twoValuesReturningFunc1() (int, error) {
	var a = 1
	return a, errors.New("error1")
}

func twoValuesReturningFunc2() (int, error) {
	var a = 2
	return a, errors.New("error2")
}
