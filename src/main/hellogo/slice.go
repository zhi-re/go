package main

import "fmt"

func main() {
	//var arrays = []int{1, 2, 3}
	var arrays = make([]int, 3, 5)
	for i := 0; i < len(arrays); i++ {
		arrays[i] = i
	}
	fmt.Println(arrays)
	fmt.Println(len(arrays))
	fmt.Println(cap(arrays))
	var nulls []int

	if nulls == nil {
		fmt.Printf("切片是空的")
	}

	/* 创建切片 */
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers)

	/* 打印原始切片 */
	fmt.Println("numbers ==", numbers)

	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Println("numbers[1:4] ==", numbers[1:4])

	/* 默认下限为 0*/
	fmt.Println("numbers[:3] ==", numbers[:3])

	/* 默认上限为 len(s)*/
	fmt.Println("numbers[4:] ==", numbers[4:])

	numbers1 := make([]int, 0, 5)
	printSlice(numbers1)

	/* 打印子切片从索引  0(包含) 到索引 2(不包含) */
	number2 := numbers[:2]
	printSlice(number2)

	/* 打印子切片从索引 2(包含) 到索引 5(不包含) */
	number3 := numbers[2:5]
	printSlice(number3)

	fmt.Println("=============================")
	var nums []int
	printSlice(nums)

	/* 允许追加空切片 */
	nums = append(nums, 0)
	printSlice(nums)

	/* 向切片添加一个元素 */
	nums = append(nums, 1)
	printSlice(nums)

	/* 同时添加多个元素 */
	nums = append(nums, 2, 3, 4)
	printSlice(nums)

	/* 创建切片 nums1 是之前切片的两倍容量*/
	nums1 := make([]int, len(nums), (cap(nums))*2)

	/* 拷贝 nums 的内容到 nums1 */
	copy(nums1, nums)
	printSlice(nums1)

}
func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
