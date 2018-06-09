package test

import "fmt"

func init() {
	fmt.Println("下划线直接执行了我，我是test1\n")
}

func Test() {
	fmt.Println("我是test函数，我不是init函数")
}
