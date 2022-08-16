package main

import (
	"fmt"
	"os"
)

type student struct {
	name string
	age  int
	id   string
}

var (
	userList map[int]*student
)

func showStudent() {
	for k, v := range userList {
		fmt.Printf("学号: %d 姓名: %s\n", k, v.name)
	}
}

func addStudent() {
	fmt.Println("addStudent")
}

func delStudent() {
	fmt.Println("popStudent")
}

func showMenu() {
	//用于实现菜单，根据用户的选择调用相关的函数
	var cmds int
	choice := `欢迎使用学生管理系统
1.查询学生
2.添加学生
3.弹出学生
4.退出
请输入你的选择(0/1/2/3): `

	for {
		fmt.Print(choice)
		_, err := fmt.Scanln(&cmds)
		if err != nil {
			return
		}
		switch cmds {
		case 1:
			showStudent()
		case 2:
			addStudent()
		case 3:
			delStudent()
		case 4:
			fmt.Println("Bye bye!")
			os.Exit(500)
		default:
			fmt.Println("Go out!")
		}
	}

}

func main() {
	userList = make(map[int]*student)

	showMenu()
}
