package main

import "fmt"

func main()  {
	fmt.Println("defer begin")

	// 将defer放入延时调用栈
	defer fmt.Println(1)

	defer fmt.Println(2)

	defer fmt.Println(3)

	fmt.Println("defer end")
}
