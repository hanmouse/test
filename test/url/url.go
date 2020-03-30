package main

import (
	"fmt"
	"net/url"
)

func main() {
	v := url.Values{"name": {"hanmouse"}, "spouse": {"jysilver"}, "son": {"jiwoo", "seungwoo"}}
	/*
		v := url.Values{}
			v.Set("name", "non-hanmouse")
			v.Set("name", "hanmouse")
			v.Add("spouse", "jysilver")
			v.Add("son", "jiwoo")
			v.Add("son", "seungwoo")
	*/

	fmt.Println(v.Get("name"))
	fmt.Println(v.Get("spouse"))
	fmt.Println(v["son"])

	fmt.Println(v.Encode())
}
