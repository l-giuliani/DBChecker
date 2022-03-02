package main

import (
	"fmt"
	libs "GLChecker/libs"
)

func main() {
	fmt.Println("test")
	config := libs.GetConfig()
	fmt.Println(config)
}