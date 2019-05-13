package main

import (
	"fmt"
	"strings"
)

func processTelnetCommand(str string, exitChan chan int) bool {

	// @close指令标识号终止本次回话
	if strings.HasPrefix(str, "@close") {
		fmt.Println("Session closed")

		// 告诉外部需要断开连接
		return false
	} else if strings.HasPrefix(str, "@shutdown") {
		fmt.Println("Server shutdown")
		// 往通道写入0， 阻塞等待接收方处理
		exitChan <- 0

		// 告诉外部断开了
		return false
	}

	fmt.Println(str)
	return true
}
