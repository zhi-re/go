package main

import "fmt"

/**
指针
*/

func main() {
	a := 10
	var aa *int
	aa = &a
	fmt.Println(*aa)

	var x, y int
	x = 10
	y = 20
	fmt.Println(x, y)
	swap(&x, &y)
	fmt.Println(x, y)
}
func swap(x *int, y *int) {
	*x, *y = *y, *x
}
