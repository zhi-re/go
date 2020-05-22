package main

import "fmt"

func main() {
	// 1----
	a := 0
	for i := 0; i <= 10; i++ {
		a += i
	}
	fmt.Println(a)

	// 2-----
	b := 1
	for ; b <= 10; {
		b += b
	}
	fmt.Println(b)

	// 3-----
	c := 1
	for c <= 10 {
		c += c
	}
	fmt.Println(c)

	sum := 0
	for {
		sum++ // 无限循环下去
	}
	fmt.Println(sum) // 无法输出
}
