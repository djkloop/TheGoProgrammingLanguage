package main

import "fmt"

// constant
const NAME = "djkloop"

// variable
var name = "the is a man"

func main() {
	/*

		go语言中的注释
	*/

	// 这种为单行注释,平常开发中使用的比较多
	/*
		这种属于多行注释
		一般有大段文字说明的时候适合使用 例如文件的功能描述或者版权描述等...
	*/

	fmt.Println("comments")
	fmt.Println(NAME)
	fmt.Println(name)

	/*

		go保留关键字
	*/


}
