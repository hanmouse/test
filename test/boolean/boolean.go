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

	var val = 0x0ff
	fmt.Printf("%#v is valid transaction id: %#v\n", val, isValidTransactionID(val))
}

func isValidTransactionID(val int) bool {
	return (0x00 <= val) && (val <= 0x0f)
}
