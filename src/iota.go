package main

import "fmt"

func main() {
	/*
		Go语言 特殊常量iota
	1.iota在const关键字出现时将被重置为0
	2.const 中每新增一行常量声明将使iota计数一次
	*/

	const a = iota
	const b = iota

	fmt.Println(a, b)

	const (
		c = iota + 2
		d = iota
		_ // 跳值
		h = 3.14 // 插队
		e = iota
	)

	fmt.Println(c, d, h, e) // 2, 1, 3.14, 4

	const (
		aa = iota * 2
		bb
		cc
	)

	fmt.Println(aa, bb, cc) // 0, 2, 4

	const (
		aaa, bbb = iota + 1, iota + 666
		ccc, ddd
		eee = iota
	)

	fmt.Println(aaa, bbb, ccc, ddd, eee)



}
