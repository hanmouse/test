package main

import (
	"encoding/json"
	"fmt"
)

// HeadphoneSystem ...
type HeadphoneSystem struct {
	Headphones []string
	DACs       []string
	AMPs       []string
}

func main() {

	headphoneSystem := HeadphoneSystem{
		Headphones: []string{"HD58X", "HD600", "HD6XX"},
		DACs:       []string{"Modi 3"},
		AMPs:       []string{"Magni Heresy", "Asgard 3"},
	}

	encoded, err := json.MarshalIndent(headphoneSystem, "", "    ")
	if err != nil {
		panic(err)
	}

	err = nil

	fmt.Println(string(encoded))

	var decoded HeadphoneSystem
	err = json.Unmarshal(encoded, &decoded)
	if err != nil {
		panic(err)
	}

	fmt.Printf("decoded: %#v\n", decoded)
}
