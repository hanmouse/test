package main

import (
	"fmt"
)

func main() {

	/*
		var tmpMap map[string]int
		tmpMap = make(map[string]int)
	*/
	/*
		tmpMap := make(map[string]int)
	*/
	tmpMap := map[string]int{}
	tmpMap["one"] = 1
	tmpMap["two"] = 2
	fmt.Printf("tmpMap=%#v\n", tmpMap)

	/*
		for {
			fmt.Printf("Hello, %#v!\n", "hanmouse")
			time.Sleep(3 * time.Second)
		}
	*/
}
