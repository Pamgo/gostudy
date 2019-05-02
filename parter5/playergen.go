package main

import "fmt"

func playerGen(name string) func() (string, int)  {

	// 血量一直为150
	hp := 150

	// 返回创建的闭包
	return func() (string, int) {
		return name,hp
	}
}

func main()  {
	generator := playerGen("high noo")

	name,ph := generator()

	fmt.Println(name,ph)
}
