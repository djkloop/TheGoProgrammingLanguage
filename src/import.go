package main

import (
	"fmt"
	"time"
	"test"
	"test2"
)

func main() {
	/*
			Go语言 基础语法 - import

		1.import 语句可以导入源代码文件所依赖的package包;
		2.不得导入源文件中没有用到的package，否则go语言编译器会报错;
		3.import语法格式
			1.导入一
				1.1 import "xxxx"
				1.2 import "yyyy"
			2.导入二
				import ( "xxxx" "yyyy" )
		4.如果一个main导入其他包,包将被顺序导入
		5.如果导入的包中依赖其他包(B), 会首先导入B,然后初始化B中常量和变量，
	    最后如果B包中有init,会自动执行init();
		6.所有包导入完成后才会对main中常量和变量进行初始化,然后执行main中的init函数(如果存在的话)
		,最后在执行main函数(有点像JavaScript同步执行代码时一行一行执行代码的感觉-_-!);
		7.如果一个包被多次导入，则只会被导入一次;
	*/
	test.Test()
	test2.Test2()
	fmt.Println("good")
	fmt.Println(time.Now().Format("2006-01-02"));
}
