package main

import (
	"example/goDemo/hello"
	"example/goDemo/variables"
	"fmt"
)

func main() {
	fmt.Println("hello world")
	fmt.Println(hello.GetName("xiaohong"))

	// go variables practice
	variables.PrcVariables()

}
