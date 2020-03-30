package myaudio_test

import (
	"hanmouse/myaudio"
	"testing"

	"github.com/bmizerany/assert"
)

func addHeadphones() *myaudio.MySystem {

	mySystem := myaudio.MySystem{}

	mySystem.AddHeadphone("HD600")
	mySystem.AddHeadphone("HD6XX")
	mySystem.AddHeadphone("HD58X")
	mySystem.AddHeadphone("KCS75")

	return &mySystem
}

func TestAddHeadphones(t *testing.T) {

	mySystem := addHeadphones()

	numHeadphone := mySystem.GetNumHeadphones()
	assert.Equal(t, 4, numHeadphone)
}

func ExampleShowHeadphones() {
	mySystem := addHeadphones()

	mySystem.ShowHeadphones()
	// Output:
	// HD600, HD6XX, HD58X, KCS75
}
