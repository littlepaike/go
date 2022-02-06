package main

import (
	"fmt"
)

//数组

//存放元素的容器
//必须指定存放的元素的类型和容量（长度
//数组长度是数组类型的一部分     如果长度不一样都不可以进行比较
func main() {
	var a1 [3]bool //[true false true]
	var a2 [4]bool //[true true false false]
	fmt.Printf("a1:%T a2:%T\n", a1, a2)

	//数组的初始化
	//如果不初始化，默认元素都是零值，布尔为false 证书和浮点型都是0，字符串：“”
	//1.初始化方式1
	a1 = [3]bool{true, true, true}
	fmt.Println(a1)
	//2.初始化方式2
	a10 := [...]int{0, 1, 2, 3} //根据初始值自动推断数组的长度是多少
	fmt.Println(a10)
	//3.初始化方式3
	a3 := [5]int{0: 1, 4: 2} //根据推断  定位赋值
	fmt.Println(a3)

	// 数组遍历
	cities := [...]string{"北京", "上海", "深圳"}
	//1.根据索引遍历
	for i := 0; i < len(cities); i++ {
		fmt.Println(cities[i])
	}
	//2. for range遍历
	for i, v := range cities {
		fmt.Println(i, v)
	}

	//多维数组
	var a11 [3][2]int
	a11 = [3][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Println(a11)

	//多维数组的遍历
	for _, v := range a11 {
		fmt.Println(v)
		for _, v2 := range v {
			fmt.Println(v2)
		}
	}

	//数组是值类型
	b1 := [3]int{1, 2, 3}
	b2 := b1
	b2[0] = 100
	fmt.Println(b1, b2)
}
