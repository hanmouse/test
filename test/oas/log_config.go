package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"
)

type LogConfigList struct {
	List []LogConfig `json:"logConfigList"`
}

type LogConfig struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func main() {

	var err error
	var rawData []byte

	configFileFromUCCMSPath := "./log_config_from_uccms.json"
	rawData, err = ioutil.ReadFile(configFileFromUCCMSPath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	logConfigList := &LogConfigList{}

	err = json.Unmarshal(rawData, logConfigList)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	logConfigMap := map[string]string{}

	for _, v := range logConfigList.List {
		logConfigMap[v.Key] = v.Value
	}

	fmt.Printf("map=%#v\n", logConfigMap)

	templateFilePath := "./logger.conf.tmpl"
	rawData, err = ioutil.ReadFile(templateFilePath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	tmpl, err := template.New("logConfigTmpl").Parse(string(rawData))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = tmpl.Execute(os.Stdout, logConfigMap)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
