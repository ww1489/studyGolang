package main

import "fmt"

/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币

写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/
var (
	coins = 100
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func dispatchCoin() (left int) {
	// 1. 依次拿到每个人的名字
	for _, uname := range users {
		// 遍历每个名字中的字符
		for _, s := range uname {
			// 根据条件分金币
			switch s {
			case 'e', 'E':
				distribution[uname]++
				coins--
			case 'i', 'I':
				distribution[uname] += 2
				coins -= 2
			case 'o', 'O':
				distribution[uname] += 3
				coins -= 3
			case 'u', 'U':
				distribution[uname] += 4
				coins -= 4
			}
		}
	}

	// 2. 拿到一个人名根据分金币的规则去分金币,
	// 2.1 每个人分的金币数应该保存到 distribution 中
	// 2.2 还要记录下剩余的金币数
	// 3. 整个第2步执行完就能得到最终每个人分的金币数和剩余金币数
	return coins
}

func main() {
	left := dispatchCoin()
	for s, i := range distribution {
		fmt.Printf("%v: %v\n", s, i)
	}
	fmt.Println("剩下：", left)
}
