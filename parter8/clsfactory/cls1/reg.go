package cls1

import (
	"fmt"
	"gostudy/parter8/clsfactory/base"
)

// 定义类
type Class1 struct {

}

func (c *Class1) Do()  {
	fmt.Println("Class1")
}

func init()  {
	// 在启动时注册类1工厂
	base.Register("Class1", func() base.Class {
		return new(Class1)
	})

}
