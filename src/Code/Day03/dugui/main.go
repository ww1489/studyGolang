package main

import "fmt"

func f(ret uint64) uint64 {
	if ret <= 1 {
		return 1
	}

	return ret * f(ret-1)
}

func main() {
	fmt.Println(f(3))
}
