package main

import "fmt"

// 值传递 不会改变原来的值
func swap(a int, b int) int {
	var temp int
	temp = a
	a = b
	b = temp
	return temp

}

// 引用传递 会改变原来的值
func swap2(a *int, b *int) {
	var temp int
	temp = *a
	*a = *b
	*b = temp

}
func main() {
	a := 1;
	b := 2;
	fmt.Println("交换之前的值：", a, b)
	swap(a, b)
	fmt.Println("交换之后的值：", a, b)
	swap2(&a, &b)
	fmt.Println("交换之后的值：", a, b)
}
