package main

import (
	"encoding/json"
	"fmt"

	_ "github.com/tidwall/pretty"
)

type JsonString []byte

func (r JsonString) String() string {
	return string(r)
}

func (r JsonString) Pretty(indent string) string {

	var js interface{}
	err := json.Unmarshal(r, &js)
	if err == nil {
		str, err := json.MarshalIndent(js, indent, "  ")
		if err == nil {
			return string(str)
		}
	}

	return string(r)
	//return string(pretty.Pretty(r))
}

// JSONMarshalIndentUser 는 JSON marshalling 시 json.Marshal() 대신 json.MarshalIndent()를 사용한다고 지정하기 위한 interface이다.
type JSONMarshalIndentUser interface {
	// 특정 type이 JSONMarshalIndentUser interface type이 되려면 해당 type에 대해 아래와 같이 빈 함수를 정의해 주면 된다.
	// func (r *SomeType) UseJSONMarshalIndent() {}
	UseJSONMarshalIndent()
}

// HeadphoneSystem ...
type HeadphoneSystem struct {
	Headphones []string `json:"headphones"`
	DACs       []string `json:"dacs"`
	AMPs       []string `json:"amps"`
}

/*
var _ JSONMarshalIndentUser = (*HeadphoneSystem)(nil)

// UseJSONMarshalIndent ...
func (r *HeadphoneSystem) UseJSONMarshalIndent() {}
*/

func main() {

	headphoneSystem := &HeadphoneSystem{
		Headphones: []string{"HD58X", "HD600", "HD6XX"},
		DACs:       []string{"Modi 3"},
		AMPs:       []string{"Magni Heresy", "Asgard 3"},
	}

	encoded, err := marshalJSON(headphoneSystem)

	err = nil

	//fmt.Println(string(encoded))
	fmt.Println(encoded.Pretty("  "))

	var decoded HeadphoneSystem
	err = json.Unmarshal(encoded, &decoded)
	if err != nil {
		panic(err)
	}

	fmt.Printf("decoded: %#v\n", decoded)
}

func marshalJSON(data interface{}) (marshalledData JsonString, err error) {
	switch data := data.(type) {
	case JSONMarshalIndentUser:
		marshalledData, err = json.MarshalIndent(data, "", "  ")
	default:
		marshalledData, err = json.Marshal(data)
	}

	return marshalledData, err
}
