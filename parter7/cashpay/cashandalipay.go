package main

import "fmt"

// 电子支付方式
type Alipay struct {

}
// 为Alipay添加CanUseFaceID()方法，表示电子支付方式刷脸
func (a *Alipay) CanUseFaceID()  {

}
// 现金支付方式
type Cash struct {

}

func (c *Cash) Stolen()  {

}

type ContainCanUseFaceID interface {
	CanUseFaceID()
}

type ContainStolen interface {
	Stolen()
}

func print(payMethod interface{})  {
	switch payMethod.(type) {
	case ContainCanUseFaceID:
		fmt.Printf("%T can use faceid\n", payMethod)
	case ContainStolen:
		fmt.Printf("%T may be stolen\n", payMethod)
	}

}

func main()  {
	print(new(Alipay))

	print(new(Cash))
}
