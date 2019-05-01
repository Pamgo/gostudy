package main

import "fmt"

// 跳出指定循环
func main()  {

	OuterLoop:
	for i := 0; i < 2; i++ {
		for j := 0; j < 5; j++ {
			switch j {
			case 2:
				fmt.Println(i,j)
				break OuterLoop
			case 3:
				fmt.Println(i,j)
				break OuterLoop
			}

		}
	}
}
