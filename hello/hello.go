package hello

import "fmt"

func GetName(name string) string {
	yourName := fmt.Sprintf("your name is %v", name)
	fmt.Println("hello world")
	fmt.Println("你好 世界")
	return yourName

}
