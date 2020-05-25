package main

import "fmt"

/**
错误处理
*/

type DivideError struct {
	dividee int
	divider int
}

func (divideError *DivideError) Error() string {
	strFormat := `
    Cannot proceed, the divider is zero.
    dividee: %d
    divider: 0`
	return fmt.Sprintf(strFormat, divideError.dividee)
}
func Divide(varDividee int, varDivider int) (result int, errorMsg string) {
	if varDivider == 0 {

		// dData := DivideError{
		//	dividee: varDividee,
		//	divider: varDivider,
		// }

		var dData DivideError
		dData.dividee = varDividee
		dData.divider = varDivider
		errorMsg = dData.Error()
		return
	} else {
		return varDividee / varDivider, ""
	}

}
func main() {
	// 正常情况
	if result, errorMsg := Divide(100, 10); errorMsg == "" {
		fmt.Println("100/10 = ", result)
	}
	// 当除数为零的时候会返回错误信息
	if _, errorMsg := Divide(100, 0); errorMsg != "" {
		fmt.Println("errorMsg is: ", errorMsg)
	}
	fmt.Println(Divide(100, 2))

	_, errorMsg2 := Divide(100, 0)
	if errorMsg2 != "" {
		fmt.Println("错误信息：", errorMsg2)
	}

}
