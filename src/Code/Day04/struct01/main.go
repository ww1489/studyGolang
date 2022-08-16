package main

import (
	"fmt"
	"os"
)

/*
1.用户操作菜单
2.用户操作函数
3.调用函数开始操作
*/

// 定义一个学生信息结构体
type student struct {
	id   int64
	name string
	age  int64
}

// 定义一个用来存储学生信息
var (
	allStudent map[int64]*student
	cmds       int
)

// 定义一个构造函数
func newStudnet(sid int64, sname string, sage int64) *student {
	return &student{
		id:   sid,
		name: sname,
		age:  sage,
	}
}

//  查询所有学生信息
func showStudent() {
	for k, v := range allStudent {
		fmt.Println("------------- User Info ---------------")
		fmt.Printf("学生ID：%d\n学生姓名：%s\n学生年纪：%d\n", k, v.name, v.age)
		fmt.Println("---------------- End ------------------")
	}
}

func addStudent() {
	var (
		uid   int64
		uname string
		uage  int64
	)
	fmt.Print("请输入学生id: ")
	fmt.Scanln(&uid)

	fmt.Print("请输入学生姓名: ")
	fmt.Scanln(&uname)

	fmt.Print("请输入学生年纪: ")
	fmt.Scanln(&uage)
	newuser := newStudnet(uid, uname, uage)
	allStudent[uid] = newuser
}

func delStudent() {
	var (
		uid int64
	)
	fmt.Print("请输删除学生id: ")
	fmt.Scanln(&uid)
	delete(allStudent, uid)
}

func showMenu() {

	choice := `欢迎使用学生管理系统！
1.查询所有学生
2.新增学生
3.删除学生
4.退出
请输入您的选项（0/1/2/3/4）: `
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
			fmt.Println("Bye Bye!")
			os.Exit(200)
		default:
			fmt.Println("GO out!")
		}

	}
}

func main() {
	// 初始化student map
	allStudent = make(map[int64]*student, 48)
	showMenu()

}
