package main

import "fmt"

type Invoker2 interface {
	Call(interface{})
}

type FuncCaller func(interface{})

func (f FuncCaller) Call(p interface{})  {
	f(p)
}

func main()  {
	var invoker Invoker2

	invoker = FuncCaller(func(v interface{}) {
		fmt.Println("from function",v)
	})

	invoker.Call("alison")
}
