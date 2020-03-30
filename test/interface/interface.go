package main

import (
	"fmt"
	"math"
	"reflect"
)

// Shape 도형의 면적, 둘레 길이 등을 구하는 method 모음
type Shape interface {
	area() float64
	perimeter() float64
}

// Rect 정의
type Rect struct {
	width, height float64
}

// Circle 정의
type Circle struct {
	radius float64
}

func (r Rect) area() float64 {
	return r.width * r.height
}

func (r Rect) perimeter() float64 {
	return (r.width + r.height) * 2
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func showAreas(shapes ...Shape) {
	for _, shape := range shapes {
		fmt.Printf("[%v] Area: %#v\n", reflect.TypeOf(shape).Name(), shape.area())
	}
}

func showPerimeters(shapes ...Shape) {
	for _, shape := range shapes {
		fmt.Printf("[%v] Perimeter: %#v\n", reflect.TypeOf(shape).Name(), shape.perimeter())
	}
}

func shapeTest() {
	rect := Rect{10, 5}
	circle := Circle{5}

	showAreas(rect, circle)
	showPerimeters(rect, circle)
}

func typeAssertionTest1() {
	var a interface{} = 1

	i := a           // a와 i 는 dynamic type, 값은 1
	j := a.(float64) // j는 int 타입, 값은 1

	fmt.Println(i) // 포인터주소 출력
	fmt.Println(j) // 1 출력
}

func typeAssertionTest2() {

	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	// 위와 같이 ok를 안 쓰면, 바로 panic이 발생한다.
	f = i.(float64) // panic
	fmt.Println(f)
}

func main() {
	//shapeTest()
	//typeAssertionTest1()
	typeAssertionTest2()
}
