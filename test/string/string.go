package main

import (
	"fmt"
	"strings"
)

func main() {
	stringerTest()
	goStringerTest()
	titleTest()
	strPtrTest()
}

// StringerType is StringerType
type StringerType string

func (r StringerType) String() string {
	return "[StringerType] " + string(r)
}

func stringerTest() {

	var str StringerType = "My type is Stringer"

	fmt.Printf("str: %s\n", str)
	fmt.Printf("str: %v\n", str)

	// "%#v"에 대해서는 GoStringer interface type이 아니기 때문에 새로 정의한 String()가 호출되지 않음.
	fmt.Printf("str: %#v\n", str)
}

// GoStringerType is GoStringerType
type GoStringerType string

// GoString is GoString
func (r GoStringerType) GoString() string {
	return "[GoStringerType] " + string(r)
}

func goStringerTest() {

	var str GoStringerType = "My type is GoStringer"

	fmt.Printf("str: %s\n", str)
	fmt.Printf("str: %v\n", str)

	// "%#v"에 대해서는 GoStringer interface type이기 때문에 새로 정의한 GoString()이 호출됨.
	fmt.Printf("str: %#v\n", str)
}

func titleTest() {
	str := "please make me a title!"
	fmt.Println(strings.Title(str))
}

func strPtrTest() {

	hello := "hello"
	ell := hello[1:4]
	hello2 := "hello"

	fmt.Printf("&hello=%#v, &hello2=%#v\n", &hello, &hello2)

	fmt.Printf("ell=%#v (ptr: %#v)\n", ell, &ell)
}
