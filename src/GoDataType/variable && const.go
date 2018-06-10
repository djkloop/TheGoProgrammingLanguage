package main

import "fmt"

func main() {
	/*
		Go语言 变量与常量

	变量声明，初始化与赋值

	*/

	// 单个变量声明和赋值
	//1.变量的声明格式：var <变量名称> [变量类型]
	//2.变量的赋值格式：<变量名称> = <值，表达式，函数等>
	//3.声明和赋值同时进行：var <变量名称> [变量类型]= <值，表达式，函数等>
	//4.分组声明格式：
	// var (
	// 	i int
	//	j float32
	//	name string
	// )
	//5.同一行声明多个变量和赋值：var a, b, c int = 1,2,3 或者 a, b := 1,2
	//6.全局变量的声明必须使用var关键字，局部可以省略
	//7.特殊变量"_"
	//8.Go语言不存在隐式转换，类型转换必须是显示的
	//9.类型转换只能发生在两种兼容的类型之间
	//10.类型转换格式：<变量名称> [:]= <目标类型>(<需要转换的变量>)

	/*

		Go语言 变量可见性规则
	大写字符开头的变量是可导出的，也就是其他包可以读取的
	小写字母开头的就是不可导出的，是私有变量
	*/

	/*

		Go  常量

	1. 常量定义也可区分为显式和隐式
	2. 常量可以使用内置表达式定义
	3.常量范围目前只支持布尔，数字，字符串
	*/

	// 显式
	// const identifier [type] = value
	const name string = "djkloop"
	fmt.Print(name)

	// 隐式
	// const identifier = value (无类型常量)
	const nName = "djkloop"
	fmt.Println("\n", nName)

	// 组合常量
	const (
		cat string = "🐱"
		dog        = "🐶"
	)

	fmt.Println("\n")
	fmt.Println(cat)
	fmt.Println("\n")
	fmt.Println(dog)

	const apple, banner string = "🍎", "🍌"
	fmt.Println(apple, banner)

	const a, c, intNumber, testBool = "🍊", "🌰", 1, true
	fmt.Println(a, c, intNumber, testBool)

	const aLen = len("string")
	fmt.Println(aLen)

}

