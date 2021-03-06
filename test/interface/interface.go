package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"reflect"
)

func main() {
	//shapeTest()
	//typeAssertionTest1()
	//typeAssertionTest2()
	//myErrorTest()
	//writerTest()
	//satisfactionTest()
	interfaceInInterfaceTest()
}

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

type myError struct {
	msg string
}

func (e *myError) Error() string {

	rval := ""

	if e != nil && e.msg != "" {
		rval = e.msg
	} else {
		rval = "Default Error"
	}

	return rval
}

func printMyError(err error) {
	fmt.Println(err.Error())
}

func myErrorTest() {
	//var myError1 = &myError{}
	var myError1 *myError = nil
	fmt.Println("type of myError1: ", reflect.TypeOf(myError1))
	printMyError(myError1)

	var myError2 = &myError{msg: "My Error!!"}
	printMyError(myError2)
}

type myWriter struct {
	filePath string
	file     *os.File
}

func (w *myWriter) openFile(filePath string) {

	var err error

	w.filePath = filePath

	w.file, err = os.OpenFile(w.filePath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
}

// io.Writer의 Write()를 구현한다.
func (w *myWriter) Write(p []byte) (n int, err error) {
	return io.WriteString(w.file, string(p))
}

func (w *myWriter) closeFile() {
	w.file.Close()
}

func writerTest() {

	w := &myWriter{}

	w.openFile("./test.txt")

	// io.Writer의 Write()를 구현한 myWriter type의 변수인 w를 인자로 입력
	fmt.Fprintf(w, "Hello!")

	w.closeFile()
}

type LegoCreator interface {
	Prepare()
	Build()
	Display()
}

type LegoCreatorImpl struct {
}

func (r LegoCreatorImpl) Prepare() {
	fmt.Println("Prepared!")
}

func (r LegoCreatorImpl) Build() {
	fmt.Println("Built!")
}

func (r LegoCreatorImpl) Display() {
	fmt.Println("Displayed!")
}

func satisfactionTest() {

	var creator LegoCreator = &LegoCreatorImpl{}

	creator.Prepare()
	creator.Build()
	creator.Display()
}

type softwareMaker interface {
	designer
	coder
	tester
}

type designer interface {
	analyzeDemand()
	writeDesignDoc()
}

type coder interface {
	analyzeDesignDoc()
	code()
}

type tester interface {
	test()
}

type devTeam struct {
	name string
}

func (r *devTeam) analyzeDemand() {
	fmt.Printf("[%s] Analyzing demand\n", r.name)
}

func (r *devTeam) writeDesignDoc() {
	fmt.Printf("[%s] Writing design document\n", r.name)
}

func (r *devTeam) analyzeDesignDoc() {
	fmt.Printf("[%s] Analyzing design document\n", r.name)
}

func (r *devTeam) code() {
	fmt.Printf("[%s] Coding\n", r.name)
}

func (r *devTeam) test() {
	fmt.Printf("[%s] Testing\n", r.name)
}

func interfaceInInterfaceTest() {

	var myTeam softwareMaker = &devTeam{name: "코어개발팀"}

	myTeam.analyzeDemand()
	myTeam.writeDesignDoc()
	myTeam.analyzeDesignDoc()
	myTeam.code()
	myTeam.test()
}
