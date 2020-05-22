package main

import "fmt"

var x, y int
var ( // 这种因式分解关键字的写法一般用于声明全局变量
	qq    int
	phone bool
)

func main() {
	fmt.Println(x, y, qq, phone)
	var name string = "aaa"
	fmt.Println(name)

	var bb, cc int = 1, 2
	fmt.Println(bb, cc)

	var age = 11
	fmt.Println(age)

	// 声明一个变量并初始化
	var a = "cq"
	fmt.Println(a)

	// 没有初始化就为零值
	var b int
	fmt.Println(b)

	// bool 零值为 false
	var c bool
	fmt.Println(c)

	userName := "ccqq"
	fmt.Println(userName)
	fmt.Println(&userName)

	// var d bool 定义的变量必须被使用
}
