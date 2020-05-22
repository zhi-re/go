package main

import "fmt"

/**

&	返回变量存储地址	&a; 将给出变量的实际地址。
*	指针变量。	*a; 是一个指针变量

 */
func main() {
	var a = 1
	var ptr *int
	ptr = &a
	fmt.Println(a)    // 1
	fmt.Println(ptr)  // 地址
	fmt.Println(*ptr) // 1
}
