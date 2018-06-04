package main

import (
	"fmt"
	"os"
	"image/gif"
	"math/rand"
)

func main() {
	// 变量
	// var name type = expression
	// 包级别的初始化在main开始之前进行，
	// 局部变量初始化和声明一样在函数执行期间进行
	var s string
	fmt.Println(s) // " "

	var i, j, k int                 // int int int
	var b, f, v = true, 2.3, "cool" // bool, float64, sting

	fmt.Println(i, j, k, b, f, v) // 0 0 0 true 2.3 "cool"

	// 变量可以通过调用返回多个值的函数进初始化
	testPath, _ := os.Getwd()
	testPath = testPath + "/T2/TestDirFile"
	var file, err = os.Open(testPath)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	fi, err := file.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, fi := range fi {
		if fi.Mode().IsRegular() {
			fmt.Println(fi.Name(), fi.Size(), "bytes")
		}
	}

	// 短变量名称
	// name := expression
	// example
	nframes := 1 // loop次数
	anim := gif.GIF{LoopCount:nframes}
	freq := rand.Float64() * 3.0
	t := 0.0
	ih := 100 // int
	var boiling float64 = 100 // float64
	var names []string
	var errs error
	x, y := 1, 2
	x, y = y, x // must type is same

	// 函数也可以用短变量来接收
	f2, err := os.Open("i wile err")
	if err != nil {
		fmt.Println(f2, err)
	}
	defer f2.Close()




	fmt.Println(boiling, names, errs, x, y, ih)
	fmt.Println(anim, freq, t) // ...只为不报错
}
