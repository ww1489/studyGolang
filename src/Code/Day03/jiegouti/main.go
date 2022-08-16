package main

import (
	"fmt"
)

// 定义一个结构体

type person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

type dog struct {
	name string
	age  int
}

func newDog(name string, age int) *dog {
	return &dog{
		name: name,
		age:  age,
	}

}

func (d dog) world() {
	fmt.Printf("%s：哈喽，小黑！\n", d.name)
}

//var susan = person{
//	name:   "susan",
//	age:    20,
//	gender: "nan",
//	hobby:  []string{"篮球", "足球", "双色球"},
//}

func shiLi(susan *person) {
	susan.name = "xman"
}

func main() {
	//// 实例化结构体
	//shiLi(&susan)
	//fmt.Println(susan.name)

	l1 := newDog("suan", 20)
	l1.world()

}
