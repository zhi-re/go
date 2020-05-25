package main

import "fmt"

/**
接口
*/

/* 定义接口 */
type Phone interface {
	call()
}

/* 定义结构体 */
type IPhone struct {
}
type Samsung struct {
}

/* 实现接口方法 */
func (IPhone IPhone) call() {
	fmt.Println("I am IPhone")

}

func (samsung Samsung) call() {
	fmt.Println("I am Samsung")

}

func main() {
	var phone Phone
	phone = new(IPhone)
	phone.call()

	phone = new(Samsung)
	phone.call()
}
