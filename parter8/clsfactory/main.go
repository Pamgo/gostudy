package main

import ("gostudy/parter8/clsfactory/base"
	_ "gostudy/parter8/clsfactory/cls1"              // 匿名引用cls1包，自动注册初始化init方法
	_ "gostudy/parter8/clsfactory/cls2"              // 匿名引用cls2包，自动注册初始化init方法
)
func main()  {
	c1 := base.Create("Class1")

	c1.Do()

	c2 := base.Create("Class2")
	c2.Do()

}
