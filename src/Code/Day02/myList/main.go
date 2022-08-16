package main

import "fmt"

// arrary练习题
// 求数组[1, 3, 5, 7, 8]所有元素的和
// 找出数组中和为指定值的两个元素的下标，比如从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0,3)和(1, 2)

func GetSum(l1 [5]int) int {
	var result int
	for i := 0; i < len(l1); i++ {
		result += l1[i]
	}
	return result
}

func GetIndex(l1 [5]int) {
	//var result string
	for i, v := range l1 {
		for i2, v2 := range l1 {
			if v+v2 == 8 {
				fmt.Printf("(%v:%v)\n", i, i2)
				//result = '(i, i2)'
			} else {
				fmt.Println("NotFount")
			}
		}
	}

}

func GetSlice() {
	s1 := []int{1, 2, 3, 4, 5}
	s1 = append(s1[:2], s1[3:]...)
	fmt.Println(s1)

}
func Point() {
	a := [3]int{1, 3, 5}
	b := a
	b[2] = 100
	fmt.Println(a, b)
}
func main() {
	//l3 := [5]int{1, 3, 5, 7, 8}
	//l2 := [5]int{1, 3, 9, 8, 11}
	//fmt.Printf("%v\n", GetSum(l3))
	//GetIndex(l2)
	Point()
	//GetSlice()
}
