package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func doPostTest() {
	reqBody := bytes.NewBufferString("Post plain text")
	resp, err := http.Post("http://httpbin.org/post", "text/plain", reqBody)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		str := string(respBody)
		println(str)
	}
}

func doPostFormTest() {
	urlValues := url.Values{
		"name":      {"hanmouse"},
		"age":       {"45"},
		"favorites": {"music", "movie"},
	}
	resp, err := http.PostForm("http://httpbin.org/post", urlValues)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		str := string(respBody)
		println(str)
	}
}

type person struct {
	Name          string   `json:"name"`
	FamilyMembers []string `json:"familyMembers"`
}

func doPostJSONTest() {
	p := person{
		Name:          "hanmouse",
		FamilyMembers: []string{"jysilver", "jiwoo", "seungwoo"},
	}
	fmt.Printf("person: %#v\n", p)

	encodedJSON, _ := json.Marshal(p)
	fmt.Printf("encodedJSON: %#v\n", string(encodedJSON))

	buff := bytes.NewBuffer(encodedJSON)
	fmt.Printf("buff: %#v\n", buff)
	resp, err := http.Post("http://httpbin.org/post", "application/json", buff)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		str := string(respBody)
		println(str)
	}
}

func main() {
	doPostTest()
	doPostFormTest()
	doPostJSONTest()
}
