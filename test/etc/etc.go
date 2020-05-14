package main

import (
	"fmt"
)

func main() {

	var str string

	str = fmt.Sprintf(":%d", 8080)

	fmt.Printf("str: %#v\n", str)

	i, s, pi := returnUninitializedReturnValues()
	fmt.Printf("i=%#v, s=%#v, pi=%#v\n", i, s, pi)

	interfaceTypeReturningFunc()

	structCopyTest()
}

func returnUninitializedReturnValues() (i int, s string, pi *int) {
	return i, s, pi
}

// interface{} 타입을 리턴한다는 건, 뭐라도 리턴해야 한다는 거다.
func interfaceTypeReturningFunc() interface{} {
	fmt.Println("interfaceTypeReturningFunc")
	return nil
}

type myProfile struct {
	id      string
	company string
}

func structCopyTest() {

	profile := myProfile{
		id:      "hanmouse",
		company: "uangel",
	}

	receiveStruct(&profile)
}

func receiveStruct(profile *myProfile) {

	var nonPointerProfile myProfile

	nonPointerProfile = *profile

	fmt.Printf("%#v\n", nonPointerProfile)
}
