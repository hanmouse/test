package main

import "fmt"

func main() {

	var str string

	str = fmt.Sprintf(":%d", 8080)

	fmt.Printf("str: %#v\n", str)

	i, s, pi := returnUninitializedReturnValues()
	fmt.Printf("i=%#v, s=%#v, pi=%#v\n", i, s, pi)
}

func returnUninitializedReturnValues() (i int, s string, pi *int) {
	return i, s, pi
}
