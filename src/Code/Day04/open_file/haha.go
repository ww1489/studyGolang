package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"

)

func openFile() {
	file, err := os.Open(`E:\test\my.cnf`)
	//file, err := os.OpenFile(`E:\test\my.cnf`, os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("Err: %v\n", err)
		return
	}
	defer file.Close()
	tmp := make([]byte, 128)
	for {
		data, err := file.Read(tmp)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("read file failed, err:", err)
			return
		}
		fmt.Printf("读取了%d字节数据\n", data)
		fmt.Println(string(tmp[:data]))
	}
}

func openFile1() {
	file, err := os.Open(`E:\test\my.cnf`)
	//file, err := os.OpenFile(`E:\test\my.cnf`, os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("Read from file failed, Err: %v\n", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("Read from file failed, Err: %v\n", err)
			return
		}
		fmt.Print(line)
	}
}

func openFile2() {
	content, err := ioutil.ReadFile(`E:\test\my.cnf`)
	if err != nil {
		fmt.Printf("Read from file failed, Err: %v\n", err)
		return
	}
	fmt.Println(string(content))
}

func main() {
	openFile2()
}
