package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

/*
 程序一般由关键字、常量、变量、运算符、类型和函数组成。

 下面列举了 Go 代码中会使用到的 25 个关键字或保留字：
 break		default		func	interface	select
  case		defer		go		map			struct
 chan		else		goto	package		switch
 const		fallthrough	if		range		type
 continue	for			import	return		var

 除了以上介绍的这些关键字，Go 语言还有 36 个预定义标识符：
 append	bool	byte	cap	close	complex	complex64	complex128	uint16
 copy	false	float32	float64		imag	int			int8		int16	uint32
 int32	int64	iota	len			make	new			nil			panic	uint64
 print	println	real	recover		string	true		uint		uint8	uintptr
*/

/*
数据类型

布尔类型、数字类型、字符串类型
派生类型:
包括：
(a) 指针类型（Pointer）
(b) 数组类型
(c) 结构化类型(struct)
(d) 联合体类型 (union)
(e) 函数类型
(f) 切片类型
(g) 接口类型（interface）
(h) Map 类型
(i) Channel 类型

数字类型
Go 也有基于架构的类型，例如：int、uint 和 uintptr。
序号		类型和描述
1		uint8
		无符号 8 位整型 (0 到 255)
2		uint16
		无符号 16 位整型 (0 到 65535)
3		uint32
		无符号 32 位整型 (0 到 4294967295)
4		uint64
		无符号 64 位整型 (0 到 18446744073709551615)
5		int8
		有符号 8 位整型 (-128 到 127)
6		int16
		有符号 16 位整型 (-32768 到 32767)
7		int32
		有符号 32 位整型 (-2147483648 到 2147483647)
8		int64
		有符号 64 位整型 (-9223372036854775808 到 9223372036854775807)
浮点型：
序号		类型和描述
1		float32
		IEEE-754 32位浮点型数
2		float64
		IEEE-754 64位浮点型数
3		complex64
		32 位实数和虚数
4		complex128
		64 位实数和虚数

以下列出了其他更多的数字类型：
序号		类型和描述
1		byte
 		类似 uint8
2		rune
 		类似 int32
3		uint
	 	32 或 64 位
4		int
 		与 uint 一样大小
5		uintptr
 		无符号整型，用于存放一个指针

数据类型		初始化默认值
int			0
float32		0
pointer		nil
*/

/*------------------ 全局变量 ------------------*/
//全局变量，全局变量允许声明但不使用。
// 全局变量可以在整个包甚至外部包（被导出后）使用。
var vbool bool
var a, b int = 10, 11
var vf float32 = 4.567
var c = "test"
var d = 88
var x, y = 109, "str"

//反引号声明的字符串就可以摆脱单行的限制,不需要使用 \ 符号避免编译器的解析错误;可以使用它来支持复杂的多行字符串
var json = `{"author": "draven", "tags": ["golang"]}`

/*------------------- 常量 --------------------*/
//常量
//常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型。
const conA, conB = "const", 99

//枚举
const (
	MON  = 1
	TUES = 2
)

//iota，特殊常量，可以认为是一个可以被编译器修改的常量。
//在每一个const关键字出现时，被重置为0，然后再下一个const出现之前，每出现一次iota，其所代表的数字会自动增加1。
const (
	ca = iota //0
	cb        //1
	cc        //2
	cd = "ha" //独立值，iota += 1
	ce        //"ha"   iota += 1
	cf = 100  //iota +=1
	cg        //100  iota +=1
	ch = iota //7,恢复计数
	ci        //8
)

//特别的
const (
	coni = 1 << iota
	conj = 3 << iota
	conk
	conl
)

/*------------------- 运算符 -------------------*/
func calc_sig() {
	var a int = 20
	var b int = 10
	var c int = 15
	var d int = 5
	var e int
	e = (a + b) * c / d // ( 30 * 15 ) / 5
	fmt.Printf("(a + b) * c / d 的值为 : %d\n", e)
	e = ((a + b) * c) / d // (30 * 15 ) / 5
	fmt.Printf("((a + b) * c) / d 的值为  : %d\n", e)
	e = (a + b) * (c / d) // (30) * (15/5)
	fmt.Printf("(a + b) * (c / d) 的值为  : %d\n", e)
	e = a + (b*c)/d //  20 + (150/5)
	fmt.Printf("a + (b * c) / d 的值为  : %d\n", e)
}

/*------------------- 循环 -------------------*/
func for_test() {
	var b int = 15
	var a int

	numbers := [6]int{1, 2, 3, 5}

	/* for 循环 */
	for a := 0; a < 10; a++ {
		fmt.Printf("a 的值为: %d\n", a)
	}

	for a < b {
		a++
		fmt.Printf("a 的值为: %d\n", a)
	}

	for i, x := range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
	}
}

/*------------------- 函数 -------------------*/
//函数返回两个数的最大值
//参数和返回值变量也是局部变量。
func max(num1, num2 int) int {
	/* 声明局部变量 */
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

//交换
func swap(x, y string) (string, string) {
	return y, x
}

/*------------------- 数组 -------------------*/
var balance1 = []float32{1000.0, 2.0, 3.4, 7.0, 50.0}
var balance2 = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

/*------------------- 指针 -------------------*/
//一个指针变量可以指向任何一个值的内存地址它指向那个值的内存地址。
func ptr_test() {
	var a int = 20 /* 声明实际变量 */
	var ip *int    /* 声明指针变量 */

	ip = &a /* 指针变量的存储地址 */

	fmt.Printf("a 变量的地址是: %x\n", &a)

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量的存储地址: %x\n", ip)

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip)

	//当一个指针被定义后没有分配到任何变量时，它的值为 nil。
	var ptr *int
	fmt.Printf("ptr 的值为 : %x\n", ptr)
}

/*------------------- 结构体 -------------------*/
//结构体是由一系列具有相同类型或不同类型的数据构成的数据集合。
//相当于java中的类
type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func struct_test() {
	var Book1 Books /* 声明 Book1 为 Books 类型 */
	var Book2 Books /* 声明 Book2 为 Books 类型 */

	/* book 1 描述 */
	Book1.title = "Go 语言"
	Book1.author = "www.w3cschool.cn"
	Book1.subject = "Go 语言教程"
	Book1.book_id = 6495407

	/* book 2 描述 */
	Book2.title = "Python 教程"
	Book2.author = "www.w3cschool.cn"
	Book2.subject = "Python 语言教程"
	Book2.book_id = 6495700

	/* 打印 Book1 信息 */
	printBook(Book1)

	/* 打印 Book2 信息 */
	printBook(Book2)
}

func printBook(book Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}

/*------------------- Range -------------------*/
// range 关键字用于for循环中迭代数组(array)、切片(slice)、链表(channel)或集合(map)的元素。
// 在数组和切片中它返回元素的索引值，在集合中返回 key-value 对的 key 值。
func range_test() {
	//这是我们使用range去求一个slice的和。使用数组跟这个很类似
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
	//在数组上使用range将传入index和值两个变量。上面那个例子我们不需要使用该元素的序号，所以我们使用空白符"_"省略了。有时侯我们确实需要知道它的索引。
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}
	//range也可以用在map的键值对上。
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	//range也可以用来枚举Unicode字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}

/*
Go语言并没有对删除切片元素提供专用的语法或者接口，需要使用切片本身的特性来删除元素。涉及到元素指针的移动。
对原始切片的增删改操作会改变底层数组，但底层数组的大小不变。
对原始切片的修改操作，会改变原始切片，但引用不变。
对原始切片的增删操作，会创建一个新的切片(引用)并返回。
因此，当业务需要大量、频繁地从一个切片中删除元素时，如果对性能要求较高的话，就需要考虑更换其他的容器了（如双链表等能快速从删除点删除元素）
*/
func sliceTest() {
	arr := [8]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("len: %d  cap: %d pointer: %p origin arr: %d\n", len(arr), cap(arr), &arr, arr)
	sli := arr[0:5] //切片的取之长度不能超过数组的长度
	fmt.Println("origin slice: ", sli)

	sli[0] = 9
	sliDel := append(sli[:2], sli[3:]...) //从角标[2]处插入，如果插入元素不等于删除元素个数，则插入位置后元素向前/后移动
	eleAdd := []int{11, 21, 31, 41, 51, 62, 71}

	arrAdded := append(eleAdd)         //创建一个新的slice并返回
	sliAdded := append(sli, eleAdd...) //在sli切片空间内扩容，对底层数组无感，返回新的切片引用

	fmt.Println("sli Del after: ", sliDel)
	fmt.Println("sli Added after: ", sliAdded)
	fmt.Println("arr Added after: ", arrAdded)
	fmt.Println("origin slice after: ", sli)
	fmt.Println("origin arr after: ", arr)
	fmt.Println("eleAddType == arrType ? :", reflect.TypeOf(eleAdd) == reflect.TypeOf(arr))
	fmt.Println("arrType == arrAddedType ? :", reflect.TypeOf(arr) == reflect.TypeOf(arrAdded))
	fmt.Println("arrAddedType == sliAddedType ? :", reflect.TypeOf(arrAdded) == reflect.TypeOf(sliAdded))
	fmt.Println("sliAddedType == sliDelType ? :", reflect.TypeOf(sliAdded) == reflect.TypeOf(sliDel))
	fmt.Println("sli == sliDel ? :", &sli == &sliDel)
}

/*------------------- Map -------------------*/
func map_test() {
	var countryCapitalMap map[string]string
	/* 创建集合 */
	countryCapitalMap = make(map[string]string)

	/* map 插入 key-value 对，各个国家对应的首都 */
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"

	/* 使用 key 输出 map 值 */
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}

	/* 查看元素在集合中是否存在 */
	captial, ok := countryCapitalMap["United States"]
	/* 如果 ok 是 true, 则存在，否则不存在 */
	if ok {
		fmt.Println("Capital of United States is", captial)
	} else {
		fmt.Println("Capital of United States is not present")
	}
}

func map_delete() {
	/* 创建 map */
	countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New Delhi"}

	fmt.Println("原始 map")

	/* 打印 map */
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}

	/* 删除元素 */
	delete(countryCapitalMap, "France")
	fmt.Println("Entry for France is deleted")

	fmt.Println("删除元素后 map")

	/* 打印 map */
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}
}

/*------------------- Channel -------------------*/
//通道（channel）是用来传递数据的一个数据结构。
//通道可用于两个 goroutine 之间通过传递一个指定类型的值来同步运行和通讯。
//操作符 <- 用于指定通道的方向，发送或接收。如果未指定方向，则为双向通道。
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 把 sum 发送到通道 c
}

//默认情况下，通道是不带缓冲区的。发送端发送数据，同时必须有接收端相应的接收数据。
func channel_no_buf() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // 从通道 c 中接收

	fmt.Println(x, y, x+y)
}

//带缓冲区的通道允许发送端的数据发送和接收端的数据获取处于异步状态，就是说发送端发送的数据可以放在缓冲区里面，可以等待接收端去获取数据，而不是立刻需要接收端去获取数据。
//不过由于缓冲区的大小是有限的，所以还是必须有接收端来接收数据的，否则缓冲区一满，数据发送端就无法再发送数据了。
//注意：如果通道不带缓冲，发送方会阻塞直到接收方从通道中接收了值。如果通道带缓冲，发送方则会阻塞直到发送的值被拷贝到缓冲区内；如果缓冲区已满，则意味着需要等待直到某个接收方获取到一个值。接收方在有值可以接收之前会一直阻塞。
func channnel_with_buf() {
	// 这里我们定义了一个可以存储整数类型的带缓冲通道
	// 缓冲区大小为2
	ch := make(chan int, 2)

	// 因为 ch 是带缓冲的通道，我们可以同时发送两个数据
	// 而不用立刻需要去同步读取数据
	ch <- 1
	ch <- 2

	// 获取这两个数据
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

//关闭通道
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func channel_close() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	// range 函数遍历每个从通道接收到的数据，因为 c 在发送完 10 个
	// 数据之后就关闭了通道，所以这里我们 range 函数在接收到 10 个数据
	// 之后就结束了。如果上面的 c 通道不关闭，那么 range 函数就不
	// 会结束，从而在接收第 11 个数据的时候就阻塞了。
	for i := range c {
		fmt.Println(i)
	}
}

/* ---------------------------- defer ---------------------------- */
//For programmers accustomed to block-level resource management from other languages, defer may seem peculiar,
//but its most interesting and powerful applications come precisely from the fact that it's not block-based but function-based.
//In the section on panic and recover we'll see another example of its possibilities.
//定义一个defer func call, 在defer作用的区域内返回时（循环、方法），立即调用执行defer函数
//资源管理
func Contents(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close() // f.Close will run when we're finished.

	var result []byte
	buf := make([]byte, 100)
	for {
		n, err := f.Read(buf[0:])
		result = append(result, buf[0:n]...) // append is discussed later.
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err // f will be closed if we return here.
		}
	}
	return string(result), nil // f will be closed if we return here.
}

//LIFO特性
//defered func的执行顺序：相当于按调用顺序将defered funcs压栈，在作用域退出返回前执行
func LIFO_test() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

//trace特性
func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}
func un(s string) {
	fmt.Println("leaving:", s)
}
func aa() {
	defer un(trace("a"))
	fmt.Println("in a")
}
func bb() {
	defer un(trace("b"))
	fmt.Println("in b")
	aa()
}

func main() {
	//局部变量，定义后必须被引用
	//使用操作符 := 可以高效地创建一个新的变量，称之为初始化声明。此种不带声明的变量只能定义在函数中
	//全局变量与局部变量声明时，名称可以相同，但是函数内的局部变量会被优先考虑。
	d, e := 8, "e"
	d = 100
	d, a = a, d

	fmt.Println("------------变量测试-------------")
	fmt.Println("hello go ...", vbool, a<<2, b>>2, vf, c, d, e)
	//值类型，直接引用内存的值；值类型的变量的值存储在栈中。
	fmt.Println("addr x: ", &x, "y: ", &y)
	//引用类型

	fmt.Println("----------------- 常量测试 --------------------")
	fmt.Println(ca, cb, cc, cd, ce, cf, cg, ch, ci)
	fmt.Println(coni, conj, conk, conl) //1 6 12 24

	fmt.Println("----------------- 运算符测试 ------------------")
	calc_sig()

	fmt.Println("----------------- 循环测试 --------------------")
	for_test()

	fmt.Println("----------------- 函数测试 --------------------")
	var ret = max(a, b)
	fmt.Printf("最大值是 : %d\n", ret)
	a, b := swap("Mahesh", "Kumar")
	fmt.Println(a, b)

	fmt.Println("----------------- 指针测试 --------------------")
	ptr_test()

	fmt.Println("----------------- 结构体测试 -------------------")
	struct_test()

	fmt.Println("----------------- slice_test -------------------")
	sliceTest()

	fmt.Println("----------------- Range_test -----------------")
	range_test()
	//循环永动机
	arr := []int{1, 2, 3}
	for _, v := range arr {
		arr = append(arr, v)
	}
	fmt.Println(arr) //1 2 3 1 2 3
	//神奇的指针
	arry := []int{1, 2, 3}
	newArr := []*int{}
	for _, v := range arry {
		newArr = append(newArr, &v) //需改为append(newArr,&arry[i])
	}
	for _, v := range newArr {
		fmt.Println(*v) //3 3 3
	}

	fmt.Println("----------------- Map_test -------------------")
	map_test()
	map_delete()

	fmt.Println("----------------- channel_test ---------------")
	channel_no_buf()
	channnel_with_buf()
	channel_close()

	fmt.Println("----------------- defer-test -----------------")
	str, err := Contents("go.mod")
	if err != nil {
		fmt.Println("read file filed!")
	} else {
		fmt.Println(str, err)
	}
	LIFO_test()
	bb()

	fmt.Println("--------------- goroutine-test ---------------")
	i := 8
	var j int
	var k int
	for j = 0; j < i; j++ {
		go fmt.Println(j, j)
	}
	fmt.Println("----")
	for k = 0; k < i; k++ {
		fmt.Println(k)
	}

}
