package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func wirteFile() {
	file, err := os.OpenFile(`E:\test\my2.cnf`, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("Read from file failed! Err: %v\n", err)
	}
	defer file.Close()
	_, err = file.Write([]byte("这是一个test string\n"))
	if err != nil {
		fmt.Printf("Write from file failed! Err: %v\n", err)
	}
	file.WriteString("这是第二个test string\n")
}

func writeFile2() {
	file, err := os.OpenFile(`E:\test\my2.cnf`, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("Read from file failed! Err: %v\n", err)
		return
	}
	defer file.Close()

	wrier := bufio.NewWriter(file)
	wrier.WriteString("这是一个测试文本： Write file for bufio")

	wrier.Flush()

}

func writeFile3() {
	str := "hello 你好"
	err := ioutil.WriteFile(`E:\test\my2.cnf`, []byte(str), 0644)
	if err != nil {
		fmt.Printf("Err: %v\n", err)
		return
	}
}
func main() {
	//wirteFile()
	writeFile2()
}
