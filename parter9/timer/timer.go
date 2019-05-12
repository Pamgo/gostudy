package main

import (
	"fmt"
	"time"
)

func main()  {
	// 创建一个打点器，没500毫秒触发一次
	ticker := time.NewTicker(time.Millisecond * 500)

	// 创建一个计时器，2秒后触发
	stopper := time.NewTimer(time.Second * 2)

	var i int

	for  {

		// 多路复用通道
		select {
		case <- stopper.C:          // 计时器到了
			fmt.Println("stop")

			goto StopHere
		case <- ticker.C:           // 打点器触发
			// 记录触发了多少次
			i ++
			fmt.Println("tick", i)
		}
	}

	StopHere:
		fmt.Println("done")
}
