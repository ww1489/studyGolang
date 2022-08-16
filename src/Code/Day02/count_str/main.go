package main

import (
	"fmt"
	"strings"
)

func countStr(s string) map[string]int {
	var k1 = make(map[string]int)
	// 切割字符串，存入列表
	s1 := strings.Split(s, " ")
	fmt.Println(s1)
	// 遍历列表，获取单词及次数 并存入map
	for _, ch := range s1 {
		k1[ch]++
	}
	// 返回结果
	return k1

}

func huiwen() {
	ss := "a上海自来水来自海上a"
	r1 := []rune(ss)

	for i := 0; i < len([]rune(ss))/2; i++ {
		if r1[i] != r1[len(r1)-1-i] {
			fmt.Println("不是回文")
			return
		}
	}
	fmt.Println("是回文")
}

func main() {
	huiwen()
}
