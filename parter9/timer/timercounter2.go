package main

import (
	"fmt"
	"runtime"
	"time"
)

func main()  {
	go func() {
		var timers int
		for  {
			timers ++
			fmt.Println("tick:",timers)
			time.Sleep(time.Second)
		}
	}()

	var input string
	fmt.Scanln(&input)

	s := runtime.NumCPU()
	fmt.Println("cpu核数：",s)
}
