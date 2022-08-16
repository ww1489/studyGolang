package main

import (
	"fmt"
	"os"
	"sort"
)

// 定义一个结构体
type Student struct {
	ID   int
	Name string
	Age  int
}

// 定义一个类结构体
type MapClass struct {
	Map map[int]*Student
}

// 定义一个类结构体构建器
func newStudnet(id int, name string, age int) *Student {
	return &Student{
		ID:   id,
		Name: name,
		Age:  age,
	}
}

// 查询学生
func (m *MapClass) showStudent() {
	// 打印表头
	fmt.Printf("%-10s%-10s%-10s\n", "ID", "name", "age")
	sortId := make([]int, 0)
	for i, _ := range m.Map {
		index := i
		sortId = append(sortId, index)
	}
	sort.Ints(sortId)
	for _, k := range sortId {
		fmt.Printf("%-10v%-10s%-10v\n", m.Map[k].ID, m.Map[k].Name, m.Map[k].Age)
	}

}

// 新增学生
func (m *MapClass) addStudent() {
	var (
		ID   int
		Name string
		Age  int
	)
	for true {
		fmt.Print("请输入学号>：")
		_, err := fmt.Scanln(&ID)
		fmt.Print("请输入姓名>：")
		_, err = fmt.Scanln(&Name)
		fmt.Print("请输入年龄>：")
		_, err = fmt.Scanln(&Age)
		if err != nil || ID == 0 || Name == "" || Age == 0 {
			fmt.Println("输入错误，重新输入！")
			continue
		}
		/*		if err != nil{
				fmt.Println("Err: err0001")
			}*/
		_, ok := m.Map[ID]
		if ok {
			fmt.Println("此学已存在")
			continue
		}
		st1 := newStudnet(ID, Name, Age)
		m.Map[ID] = st1
		break
	}
}

func (m *MapClass) delStudent() {
	var (
		ID int
	)
	for true {
		fmt.Print("请输入学号>：")
		_, err := fmt.Scanln(&ID)
		_, ok := m.Map[ID]
		if err != nil || !ok {
			fmt.Println("输入错误或者没有此学号，请重新输入！")
			continue
		}
		delete(m.Map, ID)
	}
}

func (m MapClass) showMenu() {
	//用于实现菜单，根据用户的选择调用相关的函数
	var cmds int
	choice := `欢迎使用学生管理系统
1.查询学生
2.添加学生
3.弹出学生
4.退出
请输入你的选择(1/2/3/4): `

	for {
		fmt.Print(choice)
		_, err := fmt.Scanln(&cmds)
		if err != nil {
			fmt.Println("输入错误！")
		}
		switch cmds {
		case 1:
			m.showStudent()
		case 2:
			m.addStudent()
		case 3:
			m.delStudent()
		case 4:
			fmt.Println("Bye bye!")
			os.Exit(500)
		default:
			fmt.Println("Go out!")
		}
	}

}

func main() {
	l1 := MapClass{}
	l1.Map = make(map[int]*Student, 0)
	l1.showMenu()

}
