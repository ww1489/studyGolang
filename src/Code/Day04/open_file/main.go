package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func openFile() {

	// 打开文件
	file, err := os.Open(`E:\test\my.cnf`)
	if err != nil {
		fmt.Printf("open from file failer: %v", err)
		return
	}
	defer file.Close()

	// 对文件进行读操作
	var content []byte
	tmp := make([]byte, 128)
	for true {
		data, err := file.Read(tmp)
		if data == 0 {
			break
		}
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}
		content = append(content, tmp[:data]...)
	}
	fmt.Println(string(content))
}

func readFromByBufio() {
	// 打开文件
	file, err := os.Open(`E:\test\my.cnf`)
	if err != nil {
		fmt.Println("Err:", err)
	}
	defer file.Close()
	// 读取文件
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("Err:", err)
			return
		}
		fmt.Print(line)
	}
}

func readFromByIoutil() {
	content, err := ioutil.ReadFile(`E:\test\my.cnf`)
	if err != nil {
		fmt.Println("Err:", err)
		return
	}
	fmt.Printf(string(content))
}

func main() {
	//openFile()
	//readFromByBufio()
	readFromByIoutil()

}
