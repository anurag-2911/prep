package main

import (
	"fmt"
	"goprojs/calculator"
	"goprojs/filemanager"
)

func main() {
	fmt.Println("main function")
	calculator.Calc()
	filemanager.FileManage()
}