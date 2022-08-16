package main

import "fmt"

func main() {
	slice := []int{10, 11, 12, 13}
	m := make(map[int]*int)

	for key, val := range slice {
		m[key] = &val
	}
	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k, "->", *v)
	}
}
