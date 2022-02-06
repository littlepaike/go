package main

import "fmt"

//fmt占位符
func main() {
	var n = 100
	//查看类型
	fmt.Printf("%T\n", n)
	fmt.Printf("%v\n", n)
	fmt.Printf("%b\n", n)
	fmt.Printf("%d\n", n)
	fmt.Printf("%o\n", n)
	fmt.Printf("%x\n", n)
	var s = "hello"
	fmt.Printf("字符串：%s\n", s)
	fmt.Printf("字符串：%#v\n", s)

	path := "D:\\"
	fmt.Printf(path)

	//多行的字符串
	s2 := `
		世情薄
		人情恶
		雨送黄昏花易洛
`
	fmt.Printf(s2)
}
