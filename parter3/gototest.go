package main

import (
	"fmt"
)

func main()  {

	for i := 0; i < 10; i++ {
		for y := 0; y < 10; y++ {
			if y == 2 {
				fmt.Println(y)
				goto breakhere
			}
			fmt.Println(y)
		}
	}

	// 手动返回，避免执行进入标签
	return

	breakhere:
		fmt.Println("done")
}
