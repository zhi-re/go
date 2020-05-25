package main

import "fmt"

func main() {
	/* 创建数组并传值 */
	var intArrays = []int{11, 22, 33, 44, 55}
	for i, v := range intArrays {
		fmt.Println(i, v)
	}
	/* n 是一个长度为 10 的数组 */
	var n [10]int
	for i := 0; i < len(n); i++ {
		n[i] = 100 + i
	}
	for i := 0; i < len(n); i++ {
		fmt.Println(n[i])
	}
	var floatArrays = []float32{1, 1, 1, 2}
	fmt.Println(getAverage(floatArrays, 4))
}

func getAverage(arr []float32, size float32) float32 {
	var flag, sum float32
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	flag = sum / size
	return flag

}
