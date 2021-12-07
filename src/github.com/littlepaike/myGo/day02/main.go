package main

import "fmt"

var (
	name string
	age  int
	isOk bool
)

//批量声明常量时，如果某一行声明后没有赋值，默认就和上一行一致
//const (
//	n1 = 100
//	n2
//	n3
//)
const (
	m1 = iota
	m2
	m3
	m4
)

//定义数量级
//const (
//	_  = iota
//	KB = 1 << (10 * iota)
//	MB = 1 << (10 * iota)
//	GB = 1 << (10 * iota)
//	TB = 1 << (10 * iota)
//	PB = 1 << (10 * iota)
//)

func main() {
	name = "理想"
	age = 18
	isOk = true
	//Go语言中变量声明必须使用，不然就编译不过去
	fmt.Printf("name:%s", name) //%s 占位符 使用name变量的值去替换占位符
	fmt.Println(age)
	fmt.Print(isOk)

	fmt.Println(m1)
	fmt.Println(m2)
	fmt.Println(m3)
	fmt.Println(m4)
}
