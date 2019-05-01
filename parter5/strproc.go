package main

import (
	"strings"
	"fmt"
)

// 字符串处理函数，传入字符串切片和处理链
func StringProccss(list []string, chain []func(string) string)  {

	// 遍历每一个字符串
	for index, str := range list {

		// 第一个需要处理的字符串
		result := str

		// 遍历每一个处理链
		for _,proc := range chain {
			result = proc(result)
		}

		list[index] = result
	}
}
// 自定义的移除前缀的处理函数
func removePrefix(str string) string {
	return strings.TrimPrefix(str,"go")
}

func main()  {
	// 待处理的字符串列表
	list := []string{
		"go scanner",
		"go parser",
		"go compiler",
		"go printere",
		"go farmater",
	}

	// 处理函数链
	chain := []func(string) string {
		removePrefix,
		strings.TrimSpace,
		strings.ToUpper,
	}

	StringProccss(list, chain)

	for _,str := range list{
		fmt.Println(str)
	}
}
