package main

/**
常量
 */
import (
	"fmt"
	"unsafe"
)

const (
	a = "abc"
	b = len(a)
	c = unsafe.Sizeof(a) // 实际上字符串类型对应一个结构体，该结构体有两个域，第一个域是指向该字符串的指针，第二个域是字符串的长度，每个域占8个字节，但是并不包含指针指向的字符串的内容sizeof始终返回的是16。
)

func main() {
	const length = 5
	const wide = 6
	fmt.Println(length * wide)

	fmt.Println(a, b, c)
}
