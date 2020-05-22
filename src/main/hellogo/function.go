package main

/** 函数定义
func function_name( [parameter list] ) [return_types] {
   函数体
}
 */
import "fmt"

/* 函数返回两个数的最大值 */
func max(num1, num2 int) int {
	/* 定义局部变量 */
	var result int

	if (num1 > num2) {
		result = num1
	} else {
		result = num2
	}
	return result
}

/* 函数交换*/
func swap(x, y string) (string, string) {
	return y, x
}
func main() {
	fmt.Println(max(1, 2))
	a, b := swap("cq", "12")
	fmt.Println(a, b)
}
