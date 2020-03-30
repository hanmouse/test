package hello

import (
	"camel.uangel.com/ua5g/ulib.git/uconf"
	"camel.uangel.com/ua5g/ulib.git/ulog"
)

type Hello struct {
	message string
}

func (r *Hello) Print() {
	ulog.Info("hello %s", r.message)
}

// Hello의 constructor 함수 입니다.
// constructor binding을 할 예정입니다.
func NewHello(cfg uconf.Config) *Hello {
	return &Hello{cfg.GetString("hello-message", "world")}
}
