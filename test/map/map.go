package main

import "fmt"

func main() {
	var myMap map[string]string = map[string]string{
		"one": "1",
		"two": "2",
	}

	fmt.Printf("%#v\n", myMap[""])
}
