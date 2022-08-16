# day02_作业

​	1. **如何判断一个整形数字是奇数还是偶数?**

```go
package main

import "fmt"

func main() {
	var numb1 int
	fmt.Print("请输入一个整数：")
	fmt.Scanln(&numb1)

	if numb1%2 == 1 {
		fmt.Println("奇数")
	} else {
		fmt.Println("偶数")
	}

}
```

2. **下面代码,执行会报什么错误,如何计算a+b的值**

> 答案：会报错误的运算，因为类型不匹配
>
> 要想计算要将两个变量转为同类型即可，因为200无法使用int8，所以此题a转换成int16

```go
package main

import "fmt"

func main() {
	var (
		n1 int8  = 100
		n2 int16 = 200
	)
	fmt.Println(int16(n1) + n2)
}
```

3.  **判断运算结果**

> 答案: "ls"

4.  **引导用户输入生日字符串,格式为"年-月-日",比如1990-3-16,然后按"您的生日是1990年-3月16 日的格式化字符串输出到终端控制台"**

```go
package main

import "fmt"

func main() {
	var year, mon, day int
	fmt.Print("请输入您的生日：")
	fmt.Scanf("%d-%d-%d", &year, &mon, &day)
	fmt.Printf("您的生日是%d年-%d月-%d日", year, mon, day)

}
```

5. **引导用户输入一个名字,判断是否以小写a或者大写A开头,如果是打印true,否则打印false**

```go
package main

import "fmt"

func main() {
	var n1 string
	fmt.Print("请输入名字：")
	fmt.Scanln(&n1)
	if n1[0] == 'a' || n1[0] == 'A' {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
```

6. **判断运算结果**

> 答案打印“false”

7. **var x =10 ;var y =5,x自加y如何实现**

```
var x = 10
	var y = 5
	x += y
```

8. **将整形100转为字符串100**

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var nm1 = 100
	nm2 := strconv.Itoa(nm1)
	fmt.Printf("%T\n", nm2)
}
```

9. **格式化输出函数中Print 中%T ,%#v ,%.2f 分别的作用**

> 答案: 
>
>       1. %T: 值的类型
>       1. %#v:  值的Go语法表示形式
>       1. %.2f:  浮点数保留小数点后两位

10. **在go语言中,Sprint函数和Printf函数的区别**

>答案: 

         1. Sprintf: 格式话字符串返回一个字符串,无打印
         2. printf: 格式话打印字符串，无返回值

11. **将字符串"3.24" ,转为浮点数**

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var nm1 string = "3.24"
	fnm1, err := strconv.ParseFloat(nm1, 64)
	if err != nil {
		fmt.Println("Err: ", err)
	}
	fmt.Printf("%T\n", fnm1)
}
```

12. **下面代码的执行结果是什么?并分析**

>答案: 
```
   1. a: 101
      b: 100

   2. 原因：int8为值类型，传参为拷贝，a改变不会影响bc
```


13. **某年是否为闰年,可以根据四年一闰,百年不闰,四百年闰来进行判断.也就是说,在能被4整除的年份中,除了那些能被100整除,但不能被400整除的年份外,其余的都是闰年.用户输入一个年份判断是否是闰年.**

```go
package main

import (
	"fmt"
)

func main() {
	var myear int16
	fmt.Print("请输入年份：")
	fmt.Scanln(&myear)
	if (myear%4 == 0 && myear%100 != 0) || myear%400 == 0 {
		fmt.Println(myear, "闰年")
	} else {
		fmt.Println(myear, "非闰年")
	}
}
```

14. **构建一个双位计算器:用户输入1+2或者3*4或4/2或5-8等关于加减乘除运算,打印最终的计算结果**

```go
package main

import (
	"fmt"
)

func main() {
	var (
		n1, n2 int
		s1     byte
	)
	fmt.Print("请输入公式：")
	fmt.Scanf("%d%c%d", &n1, &s1, &n2)
	switch s1 {
	case '+':
		fmt.Println(n1 + n2)
	case '-':
		fmt.Println(n1 - n2)
	case '*':
		fmt.Println(n1 * n2)
	case '/':
		fmt.Println(n1 / n2)
	}
}
```

15. **计算索引为10的斐波那契数列对应的值**

```go
package main

func f(n int) int {
	// 退出点判断
	if n == 1 || n == 2 {
		return 1
	}
	// 递归表达式
	return f(n-1) + f(n-2)
}

func main() {
	r := f(10)
	fmt.Printf("r: %v\n", r)
}
```

16. **求1+2!+3!+4!....+10!的和**

```go
package main

import "fmt"

func recursion(n1 int) int {
   if n1 == 1 {
      return 1
   } else {
      return n1 + recursion(n1-1)
   }
}

func main() {

   fmt.Println(recursion(10))

}
```

17. **程序随机内置了一个位于一定范围的数字作为猜测结果,由用户猜测此数字.用户每次猜测一次,由系统提示猜测结果:太大了,太小了,或者猜对了,直到用户猜对结果或者猜测次数用完导致失败.设定一个理想数字,比如66,让用户三次机会猜数字,如果比66大,则显示猜测的结果大了,如果比66小,则显示猜测结果小了,只有等于66,显示猜测结果正确,退出循环,最多三次都没有猜正确,退出循环,并显示"都没猜对,继续努力".**

```go
package main

import "fmt"

func main() {
	var n1 int
	for i := 0; i < 3; i++ {
		fmt.Print("请输入一个数字：")
		fmt.Scanln(&n1)
		if n1 == 66 {
			fmt.Println("猜对了")
			return
		} else if 66 < n1 {
			fmt.Println("猜大了")
		} else {
			fmt.Println("猜小了")
		}
	}
	fmt.Printf("都没猜对，正确的数字为：%d继续努力", 66)
}
```

