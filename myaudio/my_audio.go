package myaudio

import (
	"fmt"
	"strings"
)

// MySystem ...
type MySystem struct {
	headphones []string
	amps       []string
}

// AddHeadphone ...
func (sys *MySystem) AddHeadphone(name string) {

	if sys.headphones == nil {
		sys.headphones = []string{}
	}

	sys.headphones = append(sys.headphones, name)
}

// GetNumHeadphones ...
func (sys *MySystem) GetNumHeadphones() int {
	return len(sys.headphones)
}

// ShowHeadphones ...
func (sys *MySystem) ShowHeadphones() {

	var headphoneList string

	for _, headphone := range sys.headphones {
		headphoneList += (headphone + ", ")
	}

	headphoneList = strings.Trim(headphoneList, ", ")

	fmt.Println(headphoneList)
}
