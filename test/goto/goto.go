package main

import "fmt"

func main() {
	var i int
	var s string
	var p *int
	var err error

	i, s, p, err = func1()
	fmt.Println(i, s, p, err)

	i, s, p = func2()
	fmt.Println(i, s, p)
}

func func1() (i int, s string, p *int, err error) {

	err1, err2, err3 := false, true, false

	if err1 {
		i = 1
		s = "err1"
		goto returnValues
	}

	if err2 {
		i = 2
		s = "err2"
		goto returnValues
	}

	if err3 {
		i = 3
		s = "err3"
		goto returnValues
	}

returnValues:
	return i, s, p, err
}

func func2() (i int, s string, p *int) {

	err1, err2, err3 := false, false, true

	if err1 {
		i = 1
		s = "err1"
		goto returnValues
	}

	if err2 {
		i = 2
		s = "err2"
		goto returnValues
	}

	if err3 {
		i = 3
		s = "err3"
		goto returnValues
	}

returnValues:
	return i, s, p
}
