package main

/**
 *指针接收器，由于指针的特性，调用方法时，修改接收器指针的任意成员变量，在方法结束后，修改都是有效的
 */
// 定义属性结构体
type Property struct {
	value int
}

// 设置属性值
func (p *Property) SetValue(v int)  {
	// 修改p的成员变量
	p.value = v
}

// 取属性值
func (p *Property) Value() int {
	return p.value
}

func main()  {
	// 实例化属性
	p := new(Property)

	p.SetValue(100)

	fmt.Println(p.Value())
}
