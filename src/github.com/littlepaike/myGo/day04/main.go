package main

import "fmt"

//运算符

func main() {
	var (
		a = 5
		b = 2
	)
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	a++ //单独的语句,不能放到等号的右边进行赋值

	var x int
	x = 10
	x += 1
	fmt.Println(x)
}
