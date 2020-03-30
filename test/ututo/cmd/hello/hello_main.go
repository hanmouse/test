package main

import (
	"camel.uangel.com/ua5g/ulib.git/testhelper"
	"camel.uangel.com/ututo.git/hello"
	"camel.uangel.com/ututo.git/modules"
)

func main() {

	// 내가 작성한 module 목록으로부터 injector를 생성합니다.
	injector := testhelper.MainFromFile("reference.conf", modules.GetImplements())

	// 생성된 Hello 객체를 가져옵니다.
	h := injector.GetInstance((*hello.Hello)(nil)).(*hello.Hello)

	// Print를 호출하여 hello 메시지를 출력합니다.
	h.Print()
}
