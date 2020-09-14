package main

import (
	"fmt"
)

/*------------------- 隐式类型转换 -------------------*/
type TestStruct struct{}

func NilOrNot(v interface{}) bool {
	return v == nil
}

func main() {
	var s *TestStruct
	//在类型转换时，*TestStruct 类型会转换成 interface{} 类型，
	//转换后的变量不仅包含转换前的变量，还包含变量的类型信息 TestStruct，所以转换后的变量与 nil 不相等。
	fmt.Println(s == nil)    // #=> true
	fmt.Println(NilOrNot(s)) // #=> false
}
