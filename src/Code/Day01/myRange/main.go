package main

import "fmt"

//for循环range的使用
// func myRange() {
// 	s1 := "Hello 小黑！"
// 	for a, b := range s1 {
// 		fmt.Printf("%d %c\n", a, b)
// 	}

// }

//fmt占位符

// func myFmt() {
// 	var user = "susan"
// 	fmt.Printf("%+v\n", user)
// }


// 数据类型

// func myType() {
// 	//定义一个十进制变量
// 	num1 := 101

// 	//进行转换打印
// 	fmt.Printf("%b \n", num1)
// 	fmt.Printf("%o \n", num1)
// 	fmt.Printf("%x \n", num1)

// 	//定义一个八进制变量
// 	var num2 = 077

// 	//进行进制转换打印
// 	fmt.Printf("%b \n", num2)
// 	fmt.Printf("%d \n", num2)
// 	fmt.Printf("%x \n", num2)

// 	//定义一个十六进制数
// 	var(
// 		num3 = 0xff
// 	)

// 	//进行进制转换打印
// 	fmt.Printf("%b \n", num3)
// 	fmt.Printf("%d \n", num3)
// 	fmt.Printf("%x \n", num3)
// }

func myStr() {
	msg := `hello susan
hello xiaohei
hello xiaoman`

	msg2 := "c:\\nello\\tode"

	fmt.Println(msg, msg2)
	fmt.Println(msg2)
}

func main() {
	//myRange()
	//myFmt()
	//myType()
	myStr()
}
