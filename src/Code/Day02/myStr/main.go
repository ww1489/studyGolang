package main

import (
	"fmt"
	// "strings"
)

func QiePian() {
	s1 := "favourite susan"
	n1 := 2
	n2 := 4
	fmt.Printf("%c\n", s1[n1])
	fmt.Println(s1[n2])
	fmt.Println(s1[n1:n2])
}

/* func q2() {
	s := "hello world！"
	fmt.Printf("len(s): %v\n", len(s))
	fmt.Printf("strings.Split(s, \"\"): %v\n", strings.Split(s, " "))
	fmt.Printf("strings.Contains(s, \"hello\"): %v\n", strings.Contains(s, "hello"))
	fmt.Printf("strings.HasPrefix(s, \"hello\"): %v\n", strings.HasPrefix(s, "hello"))
	fmt.Printf("strings.HasSuffix(s, \"world！\"): %v\n", strings.HasSuffix(s, "world！"))
	fmt.Printf("strings.Index(s, \"l\"): %v\n", strings.Index(s, "l"))
	fmt.Printf("strings.LastIndex(s, \"l\"): %v\n", strings.LastIndex(s, "l"))
} */


func q3() {
	a := 10
	fmt.Printf("a: %b\n", a)
	b:= 3
	fmt.Printf("b: %b\n", b)

	fmt.Printf("(a ^ b): %v, %b\n", (a ^ b), (a ^ b))
	fmt.Printf("(a << 2): %v, %b\n", (a << 2), (a << 2))
	fmt.Printf("(b >> 2): %v, %b\n", (b >> 2), (b >> 2))
}

func main() {
	q3()
}
