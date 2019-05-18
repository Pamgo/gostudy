package main

import (
	"fmt"
	"runtime"
)

// 一段耗时计算
func coonsumer(ch chan int)  {

	// 无限循环获取数据的循环
	for {
		// 从通道获取数据
		data := <-ch

		if data == 0 {
			break
		}
		fmt.Println(data)
	}
}

func main()  {

	// 创建一个传递数据用的通道
	ch := make(chan int)

	for {

		var dummy string

		// 获取输入，模拟进程持续执行
		fmt.Scan(&dummy)

		if dummy == "quit"{
			for i := 0;i < runtime.NumGoroutine() -1; i++ {
				ch <- 0
			}
			continue
		}

		// 并发执行
		go coonsumer(ch)

		fmt.Println("goroutines:", runtime.NumGoroutine())
	}
}