package main

import "fmt"

func main() {
	varTrue := (1 == 1)

	if varTrue {
		fmt.Printf("varTrue is %#v\n", varTrue)
	}

	isRed := false
	if !isRed {
		fmt.Printf("It is not red\n")
	}
}
