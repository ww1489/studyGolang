package main

import "fmt"

//常量声明
//单一声明

func myConst1() {
	const user = "susan"
	const pi = 3.1415926
	const e = 2.7812
	fmt.Println(user, pi, e)
}

// 批量声明

func constMany() {
	const (
		n1 = 100
		n2
		n3
		n4 = 500
	)
	fmt.Println(n1, n3, n3, n4)
}

// iota 常量计数器

func myIota() {
	const (
		n1 = 1
		n2 = iota
		n3 = 6
	)
	fmt.Println(n1, n2, n3)
}

// 定义数量级
func mySize() {
	const (
		_  = iota
		KB = 1 << (10 * iota)
		MB = 1 << (10 * iota)
		GB = 1 << (10 * iota)
		TB = 1 << (10 * iota)
	)
	fmt.Printf("KB: %d\nMB: %d\nGB: %d\nTB: %d\n ", KB, MB, GB, TB)
}

//调用main函数
func main() {
	fmt.Println("Hello const")
	myConst1()
	constMany()
	myIota()
	mySize()
}
