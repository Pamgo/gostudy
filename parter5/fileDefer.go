package main

import (
	"os"
	"fmt"
)

func fileSize(filename string) int64 {
	// 根据文件名打开文件,返回文件句柄和错误
	f, err := os.Open(filename)

	// 如果打开发生错误，返回0
	if err != nil {
		return 0
	}
	// 取文件状态信息
	info, err := f.Stat()

	// 如果获取信息时发生错误，关闭文件并返回文件大小为0
	if err != nil {
		f.Close()
		return 0
	}

	size := info.Size()
	// 关闭文件
	f.Close()

	// 返回文件带下
	return size
}

func fileSizeBydefer(filename string) int64 {
	f, err := os.Open(filename)

	if err != nil {
		return 0
	}

	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		// defer触发，调用cllose关闭文件
		return 0
	}

	size := info.Size()

	return size
}

func main()  {
	//size := fileSize("test.ini")
	size := fileSizeBydefer("test.ini")
	fmt.Println(size)
}
