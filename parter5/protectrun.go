package main

import (
	"runtime"
	"fmt"
)

type panicContext struct {
	function string    // 所在函数
}

func ProctRun(entry func())  {

	// 延时处理函数
	defer func() {
		err := recover()

		switch err.(type) {
		case runtime.Error:   // 运行时一次
			fmt.Println("runtime error:", err)
		default:
			fmt.Println("error:", err)
		}
	}()

	entry()
}

func main()  {
	fmt.Println("运行前")

	// 允许一段手动触发错误
	ProctRun(func() {
		fmt.Println("手动宕机")

		// 使用panic传递上下文
		panic(&panicContext{
			"手动触发panic",
		})
		fmt.Println("手动宕机后")
	})

	// 故意造成空指针访问错误
	ProctRun(func() {
		fmt.Println("赋值宕机前")

		var a *int

		*a = 1

		fmt.Println("赋值宕机后")
	})

	fmt.Println("运行后")

}