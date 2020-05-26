package main

/**
Scan键盘输入录入
*/
import (
	"fmt"
	"strconv"
)

func main() {
	var start int
	var end int
	fmt.Println("请输入爬取的开始页：")
	fmt.Scan(&start)
	fmt.Println("start:", start)

	fmt.Println("请输入爬取的结束页：")
	fmt.Scan(&end)
	fmt.Println("end:", end)

	for i := start; i <= end; i++ {
		if start == 1 {
			url := "http://www.netbian.com/"
			fmt.Println(url)
		} else {
			url := "http://www.netbian.com/index_" + strconv.Itoa(i) + ".htm"
			fmt.Println(url)
		}

	}

}
