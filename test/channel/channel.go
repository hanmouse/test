package main

import (
	"fmt"
	"time"

	"github.com/otiai10/curr"
)

func main() {

	test1()
	test2()
	test3()
	test4()
}

func test1() {

	fmt.Printf("\n[%v]\n", curr.Func())

	ch := make(chan int)
	fmt.Println("[main] Created int channel")

	go func() {
		num := 100
		ch <- num
		fmt.Printf("[go routine] Sent %v to channel\n", num)
	}()

	time.Sleep(time.Second * 1)

	num := <-ch
	fmt.Printf("[main] Received %v from channel\n", num)
}

func test2() {

	fmt.Printf("\n[%v]\n", curr.Func())

	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("[goroutine] i=%#v\n", i)
			time.Sleep(time.Millisecond * 100)
		}

		// 일 다 끝냈으니 다 됐다고 메인 루틴에 알린다.
		fmt.Printf("[goroutine] printing job completed\n")
		done <- true
	}()

	fmt.Println("[main] Called goroutine for printing job")

	// 위의 Go루틴이 끝날 때까지 대기
	fmt.Println("[main] Waiting for report from goroutine")
	<-done
	fmt.Println("[main] Received job completion report from goroutine")
}

func test3() {

	fmt.Printf("\n[%v]\n", curr.Func())

	ch := make(chan string, 1)
	sendDataToChannel(ch)
	receiveDataFromChannel(ch)
}

func sendDataToChannel(ch chan<- string) {
	ch <- "Data"
	//x := <-ch // 에러 발생
}

func receiveDataFromChannel(ch <-chan string) {
	//ch <- "Data" // 에러 발생
	data := <-ch
	fmt.Println(data)
}

func test4() {

	fmt.Printf("\n[%v]\n", curr.Func())

	ch := make(chan int, 2)

	// 채널에 송신
	ch <- 1
	ch <- 2

	// 채널을 닫는다
	close(ch)

	// panic: send on closed channel
	//ch <- 3

	// 채널 수신
	/*
		println(<-ch)
		println(<-ch)

		if _, success := <-ch; !success {
			println("No more data in the channel")
		}
	*/
	for num := range ch {
		println(num)
	}
}
