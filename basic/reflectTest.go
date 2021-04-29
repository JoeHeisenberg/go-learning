package main

import (
	"fmt"
	"reflect"
)

type CustomError struct {
}

func (*CustomError) Error() string {
	return ""
}

func Add(a, b int) int { return a + b }

//由于 Go 语言的函数调用都是值传递的，所以我们只能先获取指针对应的 reflect.Value，再通过 reflect.Value.Elem 方法迂回的方式得到可以被设置的变量
func main() {
	//值传递通过反射，赋值操作导致的崩溃
	//i := 1
	//v := reflect.ValueOf(i)
	//v.SetInt(10)
	//fmt.Println(i)

	//获取指针
	j := 1
	k := reflect.ValueOf(&j) //获取变量指针
	k.Elem().SetInt(10)      //获取指针指向的变量，更新变量的值
	fmt.Println(j)

	//impl判定
	typeOfError := reflect.TypeOf((*error)(nil)).Elem()
	customErrorPtr := reflect.TypeOf(&CustomError{})
	customError := reflect.TypeOf(CustomError{})
	fmt.Println(customErrorPtr.Implements(typeOfError)) // #=> true
	fmt.Println(customError.Implements(typeOfError))    // #=> false

	//使用反射在运行期间执行方法
	v := reflect.ValueOf(Add)
	if v.Kind() != reflect.Func {
		return
	}
	t := v.Type()
	argv := make([]reflect.Value, t.NumIn())
	for i := range argv {
		if t.In(i).Kind() != reflect.Int {
			return
		}
		argv[i] = reflect.ValueOf(i)
	}
	result := v.Call(argv)
	if len(result) != 1 || result[0].Kind() != reflect.Int {
		return
	}
	fmt.Println(result[0].Int()) // #=> 1
}
