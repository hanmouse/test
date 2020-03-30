package main

import (
	"fmt"

	"github.com/savsgio/go-logger"
)

func main() {
	level := "MAJOR"
	err := fmt.Errorf("This is %v error", level)
	logger.Error(err.Error())
}
