package main

import "fmt"

// 单一变量定义

func variableProduction() {
	var name string = "susan"
	var age int = 18
	var sex bool = false
	fmt.Println(name, age, sex)
}

// 批量声明

func varibaleMany() {
	var (
		name string = "zs"
		age  int    = 17
		isOk bool   = true
	)
	fmt.Println(name, age, isOk)
}

//简短变量声明,变量推导

func variableTwo() {
	name := "susan"
	age := 19
	isOk := true
	mail, phone, user := "123@qq.com", 123567, "xaoman"
	fmt.Println(name, age, isOk)
	fmt.Printf("username: %s, mail: %s, phone: %d\n", user, mail, phone)
}

// 匿名变量
func foo() (int, string) {
	return 10, "Q1mi"
}

func variableNil() {
	x, _ := foo()
	_, y := foo()
	fmt.Println("x=", x)
	fmt.Println("y=", y)
}

func main() {
	fmt.Println("Hello Var")
	variableProduction()
	varibaleMany()
	variableTwo()
	variableNil()
}
