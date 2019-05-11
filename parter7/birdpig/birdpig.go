package main

import "fmt"

// 定义飞行动物接口
type Flyer interface {
	Fly()
}
// 定义行走动物接口
type Walker interface {
	Walk()
}
// 定义鸟类
type bird struct {

}

func (b *bird) Fly()  {
	fmt.Println("bird: fly")
}
// 为鸟添加行走方法
func (b *bird) Walk()  {
	fmt.Println("bird: walk")
}
// 定义猪
type pig struct {

}

func (p *pig) Walk()  {
	fmt.Println("pig: walk")
}

func main()  {
	animals := map[string]interface{} {
		"bird": new(bird),
		"pig": new(pig),
	}

	// 遍历映射
	for name, obj := range animals{
		// 判断对象是否为飞行动物
		f, isFly := obj.(Flyer)
		// 判断对象是否为行走动物
		w, isWalker := obj.(Walker)

		fmt.Printf("name:%s isFlayer: %v isWalker:%v\n",name,isFly,isWalker)

		if isFly {
			f.Fly()
		}
		if isWalker {
			w.Walk()
		}
	}
}