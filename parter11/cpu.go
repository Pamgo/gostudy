package main

import (
	"github.com/pkg/profile"
	"time"
)

func joinSlic() []string {

	var arr []string

	for i := 0; i < 100000; i++ {
		// 故意造成多次的切片添加操作
		arr = append(arr,"arr")
	}

	return arr
}

func joinSlice1() []string {
	const count = 10000
	var arr []string = make([]string, count)

	for i := 0; i < count; i++ {
		arr[i] = "arr"
	}
	return arr
}

func main()  {
	// 开始性能分析，返回一个停止接口
	stopper := profile.Start(profile.CPUProfile,profile.ProfilePath("."))

	// 在main()结束时停止性能分析
	defer stopper.Stop()

	// 分析的核心逻辑
	joinSlice1()

	time.Sleep(time.Second)



}
