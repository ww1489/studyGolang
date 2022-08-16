package main

import (
	"fmt"
	"strings"
)

//字符串示例类型
func stringDemo() {
	// 字符
	c1 := 'a'
	c2 := 'A'
	c3 := '中'
	//

	//字符串
	s1 := "name"
	s2 := "姓名"
	s3 := `这是一个多行字符串示例！
This is mutillinSrting Example!`

	// 转义演示
	s4 := `c:\weiyigeek\go\hello`

	//数据转义打印
	fmt.Printf("c1 char: %c, \t c2 char %c ->digtal: %d\n", c1, c2, c2)
	fmt.Println(s1, s2, c3)
	fmt.Println(s3)
	fmt.Println(s4)

	s5 := `c:\weiyigeek\go\hello`
	// 字符串常用函数（方法）
	info := "世情薄，人情恶，雨送黄昏花易落\n"
	info2 := "宋词:"
	fmt.Printf(info2 + info)

	fmt.Println("字符串分离：", strings.Split(s5, `\`))

	fmt.Println(strings.Contains(s4, "go"))
	s6 := strings.Split(s5, `\`)
	fmt.Println(strings.Join(s6, "\\"))

	fmt.Println(len(s1), len(s2))

}

func main() {
	stringDemo()
}
