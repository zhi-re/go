package main

import "fmt"

func main() {
	/* 定义局部变量 */
	var grade string = "B"
	var marks int = 90

	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}

	switch {
	case grade == "A":
		fmt.Println("优秀!\n")
	case grade == "B", grade == "C":
		fmt.Println("良好\n")
	case grade == "D":
		fmt.Println("及格\n")
	case grade == "F":
		fmt.Println("不及格\n")
	default:
		fmt.Println("差\n");
	}
	fmt.Println("你的等级是 %s\n", grade);

	fmt.Println("============")

	var x interface{}
	switch i := x.(type) {
	case nil:
		fmt.Println(" x 的类型 :%T", i)
	case int:
		fmt.Println("x 是 int 型")
	case float64:
		fmt.Println("x 是 float64 型")
	case func(int) float64:
		fmt.Println("x 是 func(int) 型")
	case bool, string:
		fmt.Println("x 是 bool 或 string 型")
	default:
		fmt.Println("未知型")
	}

	fmt.Println("==============")

	switch {
	case false:
		fmt.Println("1、case 条件语句为 false")
		fallthrough // fallthrough 会强制执行后面的 case 语句, 不会判断下一条 case 的表达式结果是否为 true。
	case true:
		fmt.Println("2、case 条件语句为 true")
		fallthrough
	case false:
		fmt.Println("3、case 条件语句为 false")
		fallthrough
	case true:
		fmt.Println("4、case 条件语句为 true")
	case false:
		fmt.Println("5、case 条件语句为 false")
		fallthrough
	default:
		fmt.Println("6、默认 case")
	}
}
