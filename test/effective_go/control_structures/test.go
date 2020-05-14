package main

import (
	"errors"
	"fmt"
)

func main() {
	redeclarationAndReassignmentTest()
	//forTest()
	//typeSwitchTest()
}

type redeclarationAndReassignment struct {
}

// 재선언(redeclaration)과 재할당(reassignment) 시험
// "v := ..." : 단축 선언
func redeclarationAndReassignmentTest() {

	r := &redeclarationAndReassignment{}

	r.test1()
	r.test2()
	r.test3()
	r.test4()
}

func (r *redeclarationAndReassignment) test1() {

	a, err := intAndErrorReturningFunc1()
	if err != nil {
		fmt.Printf("a=%v, err=%s\n", a, err)
	}

	// 위에서 a가 이미 선언되었기 때문에 a 재선언 실패
	/*
		a, err := intAndErrorReturningFunc2()
		if err != nil {
			fmt.Printf("a=%v, err=%s\n", a, err)
		}
	*/
}

func (r *redeclarationAndReassignment) test2() {

	a, err := intAndErrorReturningFunc1()
	if a == 1 && err != nil {
		// 아래 a는 위 a와 scope이 다르기 때문에 재선언 성공
		a, err := intAndErrorReturningFunc2()
		fmt.Printf("a=%v, err=%s\n", a, err)
	}
}

func (r *redeclarationAndReassignment) test3() {

	a, b, c := threeIntReturningFunc1()
	fmt.Printf("a=%v, b=%v, c=%v\n", a, b, c)

	//a, b, c := threeIntReturningFunc2()
	// 를 호출하면 a가 재선언되었다는 에러가 발생하지만,
	// 아래와 같이 새로운 변수가 하나라도 있으면, 나머지 변수들은 모두 재할당 가능
	z, b, c := threeIntReturningFunc2()
	fmt.Printf("z=%v, b=%v, c=%v\n", z, b, c)

	// 맨 앞의 변수뿐 아니라, 뒤에 변수가 새로운 변수여도 마찬가지
	a, b, d := threeIntReturningFunc2()
	fmt.Printf("a=%v, b=%v, d=%v\n", a, b, d)
}

func (r *redeclarationAndReassignment) test4() {
	// error 변수라 해도, 한 개 값만 리턴하는 함수에서는 재할당이 허용되지 않음.
	//err := errorReturningFunc()
	//err := errorReturningFunc()
}

func errorReturningFunc() error {
	return errors.New("error")
}

func intAndErrorReturningFunc1() (int, error) {
	var a = 1
	return a, errors.New("error1")
}

func intAndErrorReturningFunc2() (int, error) {
	var a = 2
	return a, errors.New("error2")
}

func twoIntReturningFunc() (int, int) {
	return 1, 2
}

func threeIntReturningFunc1() (int, int, int) {
	return 1, 2, 3
}

func threeIntReturningFunc2() (int, int, int) {
	return 10, 20, 30
}

func forTest() {

	myMap := map[string]string{
		"id":      "hanmouse",
		"company": "uangel",
	}

	// key를 생략하고자 할 땐 '_' 사용
	for _, value := range myMap {
		fmt.Println(value)
	}

	// key만 구할 때에는 "for key, _ := range ..."와 같이 할 필요 없음.
	for key := range myMap {
		fmt.Println(key)
	}
}

type typeSwitch struct {
}

func typeSwitchTest() {

	r := &typeSwitch{}

	r.test1()
}

func (r *typeSwitch) test1() {

	var t interface{}

	t = r.functionOfSomeType()

	// 여기서는 이렇게는 불가. switch 문에서만 가능.
	//t = t.(type)

	switch t := t.(type) {
	default:
		fmt.Printf("unexpected type %T\n", t) // %T prints whatever type t has
	case bool:
		fmt.Printf("boolean %t\n", t) // t has type bool
	case int:
		fmt.Printf("integer %d\n", t) // t has type int
	case *bool:
		fmt.Printf("pointer to boolean %t\n", *t) // t has type *bool
	case *int:
		fmt.Printf("pointer to integer %d\n", *t) // t has type *int
	}

	// "switch t.(type)"으로 하면, t 값에 대한 접근이 불편해짐.
	// 따라서 "switch t := t.(type)"와 같이 쓰자.
	/*
		switch t.(type) {
		default:
			fmt.Printf("unexpected type %T\n", t) // %T prints whatever type t has
		case bool:
			fmt.Printf("boolean %t\n", t) // t has type bool
		case int:
			fmt.Printf("integer %d\n", t) // t has type int
		case *bool:
			fmt.Printf("pointer to boolean %t\n", *t) // t has type *bool
		case *int:
			fmt.Printf("pointer to integer %d\n", *t) // t has type *int
		}
	*/
}

func (r *typeSwitch) functionOfSomeType() interface{} {
	var a bool = true
	return a
}
