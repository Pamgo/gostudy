package main

import "fmt"

// 调用器接口
type Invoker interface {
	// 需要实现一个Call方法
	Call(interface{})
}
// 结构体类型
type Struct struct {

}

func (s *Struct) Call(p interface{})  {
	fmt.Println("from struct",p)
}

func main()  {
	var invoker Invoker

	// 实例化结构体
	s := new(Struct)

	invoker = s

	invoker.Call("alison")

}