package main

import "fmt"

func printer(c chan int)  {
	for  {
		data := <- c

		if data == 0 {
			break
		}

		fmt.Println(data)
	}
	// 通知main已经结束循环
	c <- 0
}

func main()  {
	c := make(chan int)

	// 并发执行printer，传入channel
	go printer(c)

	for i := 1; i<=10 ; i++ {
		c <- i
	}
	// 通知并发的printer结束循环,控制printer跳出循环
	c <- 0
	// 等待printer执行完毕
	<-c
}
