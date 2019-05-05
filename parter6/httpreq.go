package main

import (
	"net/http"
	"strings"
	"fmt"
	"os"
	"io/ioutil"
)

func main()  {
	client := &http.Client{}

	// 创建一个HTTP请求
	req, err := http.NewRequest("POST","http://www.163.com/",
		strings.NewReader("key=vallue"))

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}

	// 为表头添加信息
	req.Header.Add("User-Agent","myClient")

	// 开始请求
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}

	// 读取服务器返回的内容
	data, err := ioutil.ReadAll(resp.Body)

	fmt.Println(string(data))

	defer resp.Body.Close()
}
