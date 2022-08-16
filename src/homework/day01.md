# day01_作业

1. **Golang是编译型语言还是解释性语言**

> 答案：编译型语言



2. **位和字节的关系？**

> 答案：一字节等于八个比特位



3. **uint8类型占有几个字节，取值范围是多少？**

> 答案：uint8是无符号8位整形,一个字节,取值范围0-255



4. **go build的功能是？go run的功能是？**

> 答案：go build 是将代码编译为可执行文件  会产生可执行文件
> 			go run   是编译并执行代码		 不会产生可执行文件



5. **程序中变量名命名不正确的是（多选）**

​        A.  true  B.  _ xyz   C. 2a    D. $abc

> 答案：CD



6. **声明赋值变量可以用var，比如*var* name = "yuan"，还可以不使用var的简洁声明赋值方式是?**

> 答案：name :="yuan"（此方法不适用声明全局变量)



7. **s := "hello golang"，将字符串s中的golang字符串切片出来**

> 答案：

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	// 方法1
	s := "hello golang"
	fmt.Println(s[6:])
	// 方法2
	tmp := strings.Split(s, " ")
	fmt.Println(tmp[1])

}
```



8. **将"北京 上海 广州 深圳"转换为 "北京,上海,广州,深圳"**

> 答案：

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "北京 上海 广州 深圳"
	tmp := strings.Split(s, " ")

	fmt.Println(strings.Join(tmp, ","))
}
```



9. **基于字符串操作获取用户名和密码s := "mysql ... -u root -p 123"  ，并判断用户名和密码是否是“yuan”和123，获取布尔值。**

> 答案：

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "mysql ... -u root -p 123"
	sname := strings.Contains(s, "yuan")
	spwd := strings.Contains(s, "123")
	fmt.Printf("用户名是否是”yuan“: %t, 密码是否是”123“：%t\n", sname, spwd)
}

```



10. **引导用户输入一个名字，判断改名字是否以小写a或者大写A开头，如果是打印true，否则打印false**

> 答案：

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	var uname string
	fmt.Print("请输入名字>: ")
	fmt.Scanln(&uname)
	retName := strings.HasPrefix(uname, "a") || strings.HasPrefix(uname, "A")
	fmt.Printf("判断该名字是否以小写a或者大写A开头: %t\n", retName)
}
```

 