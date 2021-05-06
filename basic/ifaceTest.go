package main

import (
	"fmt"
)

/*
struct test
*/
type error interface {
	Error() string
}

type RPCError struct {
	Code    int64
	Message string
}

func (e *RPCError) Error() string {
	return fmt.Sprintf("%s, code=%d", e.Message, e.Code)
}

func NewRPCError(code int64, msg string) error {
	return &RPCError{ // typecheck3
		Code:    code,
		Message: msg,
	}
}

func AsErr(err error) error {
	return err
}

/*
struct type & ptr type impl inface
*/
type Duck interface {
	Quack()
}

type Cat struct{}

func (c Cat) Quack() {} // 使用结构体实现接口
//func (c *Cat) Quack() {} // 使用结构体指针实现接口，不能共存

/*
隐式类型转换
*/
type TestStruct1 struct{}

func nilOrNot(v interface{}) bool {
	return v == nil
}

func main() {
	var rpcErr error = NewRPCError(400, "unknown err") // typecheck1
	err := AsErr(rpcErr)                               // typecheck2
	println(err)

	var _ Duck = Cat{}  // 使用结构体初始化变量
	var _ Duck = &Cat{} // 使用结构体指针初始化变量

	var s *TestStruct1
	fmt.Println(s == nil) // #=> true
	// 调用 NilOrNot 函数时发生了隐式的类型转换，除了向方法传入参数之外，变量的赋值也会触发隐式类型转换。
	// 在类型转换时，*TestStruct 类型会转换成 interface{} 类型，转换后的变量不仅包含转换前的变量，还包含变量的类型信息 TestStruct，所以转换后的变量与 nil 不相等。
	fmt.Println(nilOrNot(s)) // #=> false
}
