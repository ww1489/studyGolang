package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func openFile() {
	// 打开文件句柄
	file, err := os.Open(`E:\test\my.cnf`)
	if err != nil {
		fmt.Println("Read from file failed, Err：", err)
		return
	}
	defer file.Close()

	// 使用for循环读取文件
	tmp := make([]byte, 128)
	content := []byte{}
	for {
		data, err := file.Read(tmp)
		if err == io.EOF {
			fmt.Printf("Read from file failed, Err：%s", err)
			break
		}
		if err != nil {
			fmt.Printf("Read from file failed, Err：%s", err)
			return
		}
		content = append(content, tmp[:data]...)
	}
	fmt.Println(content)

}

func openFile2() {
	// 使用os.open打开文件
	file, err := os.Open(`E:\test\my.cnf`)
	if err != nil {
		fmt.Println("Read from file failed, Err：", err)
		return
	}
	defer file.Close()
	// 使用bufio读取文件
	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("Read from file failed, Err：", err)
			return
		}
		if err != nil {
			fmt.Println("Read from file failed, Err：", err)
			return
		}
		fmt.Print(line)
	}
}

func openFile3() {
	date, err := ioutil.ReadFile(`E:\test\my2.cnf`)
	if err != nil {
		fmt.Println("Read from file failed, Err：", err)
		return
	}
	fmt.Println(string(date))
}

func writeFile() {
	file, err := os.OpenFile(`E:\test\my2.cnf`, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Read from file failed, Err：", err)
		return
	}
	defer file.Close()
	lens, err := file.Write([]byte("这是一个测试文本 for Writefile"))
	if err != nil {
		fmt.Println("Write file dailed, Err: ", err)
		return
	}
	fmt.Printf("写入成功！写入了%d个字节", lens)
}

func writeFileFromBufio() {
	file, err := os.OpenFile(`E:\test\my2.cnf`, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Read from file failed, Err：", err)
		return
	}
	defer file.Close()

	// 使用bufio.NewWrite方法写入文件
	writer := bufio.NewWriter(file)
	lens, err := writer.WriteString("这是一个测试文本 for Bufio")
	if err != nil {
		fmt.Println("Write file failed, Err: ", err)
		return
	}
	err = writer.Flush()
	if err != nil {
		fmt.Println("Write file failed, Err: ", err)
		return
	}
	fmt.Printf("写入成功！写入了%d个字节", lens)
}

func writeFileFromIoutil() {
	err := ioutil.WriteFile(`E:\test\my3.cnf`, []byte("测试文本 ioutil"), 0644)
	if err != nil {
		fmt.Println("err")
		return
	}
}

func main() {
	openFile()
	//openFile2()
	//writeFile()
	//writeFileFromBufio()
	//writeFileFromIoutil()
}
