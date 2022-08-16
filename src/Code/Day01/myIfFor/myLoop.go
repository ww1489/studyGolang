package main

import "fmt"

/*func myIf() {
	n1 := 5
	if n1 < 10 {
		fmt.Println("hahahah")
	} else if n1 > 3 {
		fmt.Println("xixiiix")
	} else {
		fmt.Println("llll")
	}
}

func myRange() {
	s1 := "hello world 小go"

	for a, b := range s1 {
		fmt.Printf("%v %c \n ", a, b)

	}
}
*/
func main() {
	n1 := 5
	for ; n1 < 10; n1++ {
		fmt.Printf("这是第%d次循环\n", n1)
	}
	n2 := 0x4E2D
	fmt.Printf("%c", n2)
}
