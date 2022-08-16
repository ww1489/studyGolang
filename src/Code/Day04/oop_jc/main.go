package main

import (
	"fmt"
)

type Person struct {
	id    int
	name  string
	age   int
	hobby []string
}

func newPerson(id int, name string, age int, hobby []string) *Person {
	return &Person{
		id:    id,
		name:  name,
		age:   age,
		hobby: hobby,
	}
}

func main() {
	l1 := newPerson(001, "susan", 23, []string{"sing", "jump", "rap", "basketball"})
	fmt.Printf("%v\n", l1)
}
