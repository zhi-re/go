package main

/**
闭包是匿名函数，可在动态编程中使用
*/
import "fmt"

func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}

}

func main() {
	/* nextNumber 为一个函数，函数 i 为 0 */
	nextNumber := getSequence()
	/* 调用 nextNumber 函数，i 变量自增 1 并返回 */
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
}
