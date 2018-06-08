package main

import (
	"fmt"
	"strings"
	"os"
)

func main() {
	// for循环几种写法
	//for initialization; condition; post {
	//   doSomething
	//}

	// initialization -> 可选 for循环之前会优先处理initialization逻辑
	// condition -> 可选 必须是一个结果为boolean的值表达式,每次循环都会检查是否满足这个boolean 当boolean为false则循环结束
	// post -> 可选 迭代结束之后被执行

	// same JavaScript while loop
	// for condition {
	// 	  doSomething
	// }

	// infinite loops 不过可以通过break或return语句来结束循环
	// for {
	//   doSomething
	// }

	// _为go语言的空白标识符
	//s, sep := "", ""
	//for _, arg := range os.Args[1:] {
	//	s += sep + arg
	//	sep = " "

	// 使用strings函数优化下
	/*
		 ~/go/TheGoProgrammingLanguage/T1/ [master+*] go run ./T1-for.go hello build good kkk
		hello build good kkk
	*/
	fmt.Println(strings.Join(os.Args[1:], " ")) // hello build good kkk
}
