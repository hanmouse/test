package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-akka/configuration"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Printf("Usage: %v <config_file_path>\n", filepath.Base(os.Args[0]))
		return
	}

	configFilePath := os.Args[1]

	conf := configuration.LoadConfig(configFilePath)

	draList := conf.GetValue("DRA_LIST").GetArray()

	for _, dra := range draList {
		obj := dra.GetObject()
		host := obj.GetKey("host").GetString()
		ipaddr := obj.GetKey("ipaddr").GetString()
		fmt.Println(host, ipaddr)
	}
}
