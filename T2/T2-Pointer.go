package main

import "fmt"

func main() {
	xx := 1
	p := &xx        // p 是整型指针,指向xx
	fmt.Println(*p) // 1
	*p = 2          // change xx pointer
	fmt.Println(xx) // 2
	var x, y int
	fmt.Println(&x == &x, &x == &y, &x == nil, &y == nil) // true, false, false, false

	// 函数返回局部变量的地址
	var pF = f()
	fmt.Println(pF) // 这里指针包含的是一个变量地址

	//
	n := 1
	incr(&n)              // 2
	fmt.Println(incr(&n)) // 3
}

func incr(p *int) int {
	*p++ // 递增p所指向的数值, p自身保持不变
	return *p
}

func f() *int {
	v := 1
	return &v
}
