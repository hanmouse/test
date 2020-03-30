package modules

import (
	"camel.uangel.com/ututo.git/hello"
	"github.com/csgura/di"
)

type HelloModule struct {
}

func (r *HelloModule) Configure(binder *di.Binder) {
	// Hello 타입과 constructor 함수를 bind 하였습니다.
	// Hello 타입에 대한 GetInstance 요청이 있으면 NewHello 를 실행하여 Hello 객체를 생성할 것입니다.
	// NewHello 함수에 필요한 arguement는 injector에 의해서 자동으로 생성 됩니다. ( bind 되어 있는 경우만 )
	binder.BindConstructor((*hello.Hello)(nil), hello.NewHello)
}
