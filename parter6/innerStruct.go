package main

import "fmt"

func printMsg(msg *struct{
	id int
	data string
}){
	fmt.Println("%T\n",msg)
}

func main()  {
	msg := &struct {
		id int
		data string
	}{
		1024,
		"hello",
	}

	printMsg(msg)
}