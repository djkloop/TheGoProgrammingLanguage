package main

import (
	"fmt"
	"unsafe"
	"reflect"
)

type Go语言学习 int32

func main() {
	/*
			Go语言 数据类型
		整型，浮点型，复数，字符串和布尔型
		派生类型
	*/

	// 整型
	var i uint8 = 1
	fmt.Println(unsafe.Sizeof(i)) // 1 byte = 8bit

	var ii int32 = 1
	fmt.Println(unsafe.Sizeof(ii)) // 4 byte = 32bit

	var i8 int = 1
	fmt.Println(unsafe.Sizeof(i8)) // 根据当前系统64位为8 32位系统为4

	var iNumberF float32 = 1.2
	fmt.Println(iNumberF) // 1.2  float必须是32或64，必须带数字不能像整型那样

	var iNumberFTwo float64 = 2.3
	fmt.Println(iNumberFTwo) // 2.3

	var trueBool bool = true;
	fmt.Println(trueBool) // true 只能是bool类型不能是其他类型

	var falseBool bool = false;
	fmt.Println(falseBool) // false

	var iNumber byte = 1
	fmt.Println(iNumber) // 1 byte

	var runeNumber rune = 1
	fmt.Println(unsafe.Sizeof(runeNumber)) // 4

	/*

			Go语言类型零值和类型别名
		类型零值不是空值，而是某个变量被声明后的默认值，一般情况下，值类型的默认值为0，布尔类型的为false，
		string默认值是空串
	*/

	var defI int
	var defI1 int32
	var defI2 int64
	var	defF32 float32
	var defF64 float64
	var bT bool
	var defC complex64

	fmt.Println(defI, defI1, defI2, defF32, defF64, defC, bT)

	var defS string
	fmt.Printf("%q \n", defS)


	/*

		Go语言别名
	*/


	var q Go语言学习
	var qq Go语言学习 = 1
	fmt.Print(q, reflect.TypeOf(q))
	fmt.Print( "\n", q + qq) // 1
}
