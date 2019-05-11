package main

import (
	"fmt"
	"io"
)

// 定义一个设备
type device struct {

}
// 实现io.Writer的Write()方法
func (d *device) Write(p []byte) (n int, err error) {
	fmt.Println(string(p))
	return 0,nil
}
func (d *device) Close()  error {
	return nil
}

func main()  {
	// 声明写入关闭器，并赋予device的实例
	var wc io.WriteCloser = new(device)

	// 写入数据
	wc.Write(nil)

	// 关闭设备
	wc.Close()

	// 声明写入器，并赋予device的新实例
	var writeOnly io.Writer = new(device)

	// 写入数据
	writeOnly.Write(nil)
}
