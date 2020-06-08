package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Now()
	fmt.Println(t.Format("20060102150405"))

	//当前时间戳
	t1 := time.Now().Unix()
	fmt.Println(t1)
	//时间戳转化为具体时间
	fmt.Println(time.Unix(t1, 0).String())

	//基本格式化的时间表示
	fmt.Println(time.Now().String())

	fmt.Println(time.Now().Format("2006year 01month 02day"))

	timeUnix := time.Now().Unix()
	formatTimeStr := time.Unix(timeUnix, 0).Format("2006-01-02")
	var ss string
	ss = formatTimeStr
	fmt.Println(ss)
}
