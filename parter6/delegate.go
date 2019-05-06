package main

import "fmt"

type class struct {

}

func (c *class) Do(v int)  {
	fmt.Println("call method do:", v)
}

func funcDo(v int)  {
	fmt.Println("call function do:", v)
}

func main()  {

	// 声明一个函数回调
	var delegate func(int)
	// 创建结构体实例
	c := new(class)
	// 将回调设为c的Do方法
	delegate = c.Do
	// 调用
	delegate(100)

	// 将回调设置为普通函数
	delegate = funcDo

	delegate(100)


}
