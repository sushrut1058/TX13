package main

import (
	"fmt"
	"promptgen/classification"
	"promptgen/core"
	"promptgen/types"
)

func main() {
	// path := "../Reverse-Proxy"
	path := "../open-source/awx"
	projectType := classification.GetType(path)
	switch projectType {
	case types.Django:
		fmt.Println("It's django, handling...")
	default:
		fmt.Println("Not supported yet")
	}
	fmt.Println(core.FindUrlsPy(path))
}

//make tree
//give
