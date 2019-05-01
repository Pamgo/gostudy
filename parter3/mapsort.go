package main

import (
	"sort"
	"fmt"
)

func main()  {
	scene := make(map[string] int)

	scene["route"] = 66
	scene["alison"] = 99
	scene["china"] = 100

	// 声明一个切片保存map数据
	var sceneList []string

	// 将map数据遍历到切片中
	for k := range scene{
		sceneList = append(sceneList,k)
	}

	// 对切片进行排序
	sort.Strings(sceneList)

	// 输出
	fmt.Println(sceneList)
}
