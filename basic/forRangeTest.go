package main

import "fmt"

func main() {
	/*
		循环次数验证
	*/
	arr := []int{1, 2, 3}
	for _, v := range arr {
		arr = append(arr, v)
	}
	fmt.Println(arr) //1 2 3 1 2 3

	/*
		不恰当的引用
		对于所有的 range 循环，Go 语言都会在编译期将原切片或者数组赋值给一个新的变量 ha，在赋值的过程中就发生了拷贝，所以我们遍历的切片已经不是原始的切片变量了。
		而遇到这种同时遍历索引和元素的 range 循环时，Go 语言会额外创建一个新的 v2 变量存储切片中的元素，循环中使用的这个变量 v2 会在每一次迭代被重新赋值，在赋值时也发生了拷贝。
		因为在循环中获取返回变量的地址都完全相同，所以最终都指向同一个值
	*/
	arr1 := []int{1, 2, 3}
	newArr := []*int{}
	for _, v2 := range arr1 {
		newArr = append(newArr, &v2)
	}
	for _, v2 := range newArr {
		fmt.Println(*v2) //3 3 3
	}

	/*
		使用元素地址
	*/
	arr2 := []int{1, 2, 3}
	newArr2 := []*int{}
	for i, _ := range arr2 {
		newArr2 = append(newArr2, &arr2[i])
	}
	for _, v := range newArr2 {
		fmt.Println(*v) //1 2 3
	}

	/*
		map的随即遍历特性
	*/
	hash := map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
	}
	for k, v := range hash {
		println(k, v)
	}
}
