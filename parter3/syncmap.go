package main

import (
	"sync"
	"fmt"
)

func main()  {
	var scene sync.Map

	// 将键值对保存到sync.map中
	scene.Store("greece",97)
	scene.Store("alison",100)
	scene.Store("china",999)

	// 从sync.map中取值
	fmt.Println(scene.Load("alison"))

	// 根据键删除对应的键值对
	scene.Delete("china")
	// 遍历所有sync.map中的键值对
	scene.Range(func(key, value interface{}) bool {
		fmt.Println("iterate:",key,value)
		return true
	})
}
