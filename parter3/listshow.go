package main

import (
	"container/list"
	"fmt"
)

func main()  {
	l := list.New()

	// 尾部添加
	l.PushBack("alison")

	// 头部添加
	l.PushFront("china")

	element := l.PushBack("first")
	// 在first之后添加
	l.InsertAfter("high", element)
	// 在first之前添加
	l.InsertBefore("noon",element)

	l.Remove(element)
	// 遍历list

	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}
