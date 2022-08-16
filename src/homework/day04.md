# day04_作业



1. 代码执行结果：

```
const (
     x = iota
     _
     y
     z = "zzz"
     k 
     p = iota
 )

func main()  {
    fmt.Println(x,y,z,k,p)
}
```

执行结果：

```go
0 2 zzz zzz 5
```



2. 代码执行结果：

```go
slice := []int{10, 11, 12, 13}
m := make(map[int]*int)

for key, val := range slice {
   m[key] = &val
}
fmt.Println(m)

for k, v := range m {
   fmt.Println(k, "->", *v)
}
```

执行结果：

```go
map[0:0xc0000140a8 1:0xc0000140a8 2:0xc0000140a8 3:0xc0000140a8]
2 -> 13
3 -> 13
0 -> 13
1 -> 13
```



3. 基于切片和map数据类型实现一个客户信息管理系统

```go
1. 模拟实现基于⽂本界⾯的《客户关系管理软件》；
2. 该软件能够实现客户对象（map存储）向容器对象（切片）的插⼊，修改和删除，并且能够打印客户信息明细表。
3. 项⽬的界⾯设计：参考以下图片
```

```go
package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

// 定义一个value为[]interface{}类型的map
var (
	Customers  = make(map[int][]interface{})
	CustomerId int
)

// 获取用户信息
func getCustomerInfo() (string, string, int16, string) {
	var (
		name, gender, email string
		age                 int16
	)
	fmt.Print("请输入客户姓名>：")
	fmt.Scan(&name)

	fmt.Print("请输入客户性别>：")
	fmt.Scan(&gender)

	fmt.Print("请输入客户年龄>：")
	fmt.Scan(&age)

	fmt.Print("请输入客户邮箱>：")
	fmt.Scan(&email)

	return name, gender, age, email
}

// 新增一个用户
func addCustomer() {
	fmt.Printf("\033[1;35;40m%s\033[0m\n", "------------------------- 新增用户 ---------------------------")
	name, gender, age, email := getCustomerInfo()
	CustomerId++
	newCustomer := []interface{}{name, gender, age, email}
	Customers[CustomerId] = newCustomer
	fmt.Printf("\033[1;35;40m%s\033[0m\n", "用户添加完成！")
}

// 删除一个用户
func delCustomer() {
	var delCustomerId int
	fmt.Printf("\033[1;31;40m%s\033[0m\n", "------------------------- 删除用户 --------------------------")
	fmt.Print("请输入需要删除的用户id>：")
	fmt.Scan(&delCustomerId)
	_, ok := Customers[delCustomerId]
	if !ok {
		fmt.Printf("\033[1;31;40m%s\033[0m\n", "用户id不存在！")
		return
	}
	delete(Customers, delCustomerId)
	fmt.Printf("\033[1;31;40m%s\033[0m\n", "用户删除完成！")
}

// 列出所有用户
func listCustomer() {
	fmt.Printf("\033[1;32;40m%s\033[0m\n", "------------------------- 查询用户 --------------------------")
	var tmp []int
	for i := range Customers {
		tmp = append(tmp, i)
	}
	// 对key进行排序
	sort.Ints(tmp)
	// 根据排序取出值
	for _, v := range tmp {
		fmt.Printf("编号：%-10d姓名：%-10s性别：%-10s年龄：%-10d邮箱：%-10s\n", v, Customers[v][0], Customers[v][1],
			Customers[v][2], Customers[v][3])
	}
	fmt.Printf("\033[1;32;40m%s\033[0m\n", "用户查询完成！")
}

// 修改用户信息
func editCustomer() {
	var editCustomerId int
	fmt.Printf("\033[1;36;40m%s\033[0m\n", "------------------------- 修改用户 --------------------------")
	fmt.Print("请输入需要修改的用户id>：")
	fmt.Scan(&editCustomerId)
	_, ok := Customers[editCustomerId]
	if !ok {
		fmt.Printf("\033[1;36;40m%s\033[0m\n", "用户id不存在")
		return
	}
	name, gender, age, email := getCustomerInfo()
	newCustomer := []interface{}{name, gender, age, email}
	Customers[editCustomerId] = newCustomer
	fmt.Printf("\033[1;36;40m%s\033[0m\n", "用户修改完成")
}

// 判断是否返回上一层
func isBack() bool {
	// 引导用户选择继续还是返回
	fmt.Print("请问是否继续，N返回上一层【Y/N】:")
	var backCmd string
	fmt.Scan(&backCmd)
	if strings.ToUpper(backCmd) == "Y" {
		return true
	} else {
		return false
	}
}

// 展示菜单
func showMenu() {
	var cmd int
	choice := `----------------- 欢迎使用客户信息管理系统 ------------------
1.查询客户信息
2.新增客户信息
3.删除客户信息
4.修改客户信息
0.退出
请输入您的选项（0/1/2/3/4）: `
	for {
		fmt.Printf("\033[1;33;40m%s\033[0m", choice)
		_, err := fmt.Scan(&cmd)
		if err != nil {
			fmt.Println("请输入正确的选项!")
			continue
		}
		switch cmd {
		case 1:
			listCustomer()
		case 2:
			for {
				addCustomer()
				ret := isBack()
				if ret {
					continue
				}
				break
			}
		case 3:
			for {
				delCustomer()
				ret := isBack()
				if ret {
					continue
				}
				break
			}
		case 4:
			for {
				editCustomer()
				ret := isBack()
				if ret {
					continue
				}
				break
			}
		case 0:
			fmt.Println("Bye Bye!")
			os.Exit(0)
		default:
			fmt.Println("请输入正确的选项!")
		}
	}
}

func main() {
	showMenu()
}

```

