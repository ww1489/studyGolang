package main

import (
	"fmt"
)

type AllMover interface {
	Move()
}

// 定义一Dog个结构体
type Dog struct {
	Family string
	Name   string
	Foot   int8
}

// 定义一个Cat结构体
type Insect struct {
	Family string
	Name   string
	Foot   int8
}

/*// 定义一个构造函数
func newDog(family, name string, foot int8) *Dog {
	return &Dog{
		Family: family,
		Name:   name,
		Foot:   foot,
	}
}

func newInsect(family, name string, foot int8) *Insect {
	return &Insect{
		Family: family,
		Name:   name,
		Foot:   foot,
	}
}*/

// 定义两个方法
func (a Dog) Move() {
	fmt.Printf("%s有%d条腿\n%s正在移动\n", a.Name, a.Foot, a.Name)
}

func (i Insect) Move() {
	fmt.Printf("%s有%d条腿\n%s正在移动\n", i.Name, i.Foot, i.Name)
}

func da(x AllMover) {
	x.Move()
}

func main() {
	dahang := Dog{
		Family: "中华田园犬",
		Name:   "大黄",
		Foot:   4,
	}
	xiaohuang := Insect{
		Family: "大黄峰",
		Name:   "小黄",
		Foot:   6,
	}
	// 调用方法
	da(dahang)
	da(xiaohuang)

	var ss AllMover
	ss = dahang
	ss.Move()

}
