package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main()  {

	// 声明一个等待组
	var wg sync.WaitGroup

	// 准备一些列网址
	var urls = []string{
		"http://www.github.com/",
		"http://www.baidu.com/",
		"https://www.golangtc.com/",
	}

	for _, url := range urls{
		// 每一个任务开始时，将等待组增加1
		wg.Add(1)

		go func(url string) {

			// 使用defer，表示函数完成时将等待组值减1
			defer wg.Done()

			// 使用http访问提供的地址
			_, err := http.Get(url)

			fmt.Println(url, err)
		}(url)

		wg.Wait()

		fmt.Println("over")
	}
}
