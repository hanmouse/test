package modules

import (
	ulibmodule "camel.uangel.com/ua5g/ulib.git/modules"
	"github.com/csgura/di"
)

func GetImplements() *di.Implements {
	impls := di.NewImplements()

	// HelloModule을 추가했습니다.
	impls.AddImplement("HelloModule", &HelloModule{})

	// ulib에서 제공하는 module도 같이 등록합니다.
	impls.AddImplements(ulibmodule.GetImplements())
	return impls
}
