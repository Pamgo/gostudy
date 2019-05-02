package main

import "fmt"

func main()  {
	str := "hello world"

	foo := func() {
		str = "hello alison"
	}

	foo()
	fmt.Println(str)
}
