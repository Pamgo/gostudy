package main

import (
	"fmt"
	"time"
)

func main()  {
	exit := make(chan int)

	fmt.Println("start")

	time.AfterFunc(time.Second, func() {
		fmt.Println("on second after")

		// 通知main()的goroutine已经结束
		exit <- 0
	})
	// 等待结束
	<- exit
}
