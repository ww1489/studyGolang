package main

import (
	"fmt"
	"io"
	"os"
)

func CopyFile(srcfile, dstfile string) (winer int64, err error) {
	src, err := os.Open(srcfile)
	if err != nil {
		fmt.Printf("Open file %s error, Err: %v\n", srcfile, err)
	}
	defer src.Close()

	dst, err := os.OpenFile(dstfile, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("Write file %s error, Err: %v\n", dstfile, err)
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}

func main() {
	src := `E:\test\picture.jpg`
	dst := `E:\test\picture2.jpg`
	_, err := CopyFile(src, dst)
	if err != nil {
		fmt.Printf("Copy file error, Err: %v\n", err)
		return
	}
	fmt.Println("Copy done!")
}
