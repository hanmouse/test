package main

import (
	"fmt"

	"github.com/go-akka/configuration"
)

var configText = `
DRA_LIST = [
    {
    	host = JKT1
        ipaddr = 10.37.202.77
        dbname = udra
        port = 3306
        user = root
        pass = root.1231
    }
  {
    	host = JKT2
        ipaddr = 10.41.113.1
        dbname = udra
        port = 3306
        user = root
        pass = root.123
    }
]
`

func main() {

	conf := configuration.ParseString(configText)

	draList := conf.GetValue("DRA_LIST")

	list := draList.GetArray()
	for _, dra := range list {
		obj := dra.GetObject()
		host := obj.GetKey("host").GetString()
		ipaddr := obj.GetKey("ipaddr").GetString()
		fmt.Println(host, ipaddr)
	}
}
