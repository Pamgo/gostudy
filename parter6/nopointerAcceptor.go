package main

import "fmt"

type Point struct {
	X int
	Y int
}

// 非指针接收器的加法
func (p Point) Add(other Point) Point {
	// 成员值与参数相加后返回新的结构
	return Point{p.X + other.X,p.Y + other.Y}
}

func main()  {
	// 初始化点
	p1 := Point{1,1}
	p2 := Point{2,3}

	result := p1.Add(p2)

	fmt.Println(result)
}