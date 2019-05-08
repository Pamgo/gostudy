package main

import "fmt"

// 一个服务需要满足能够开启和写日志的功能
type Service interface {
	Start()             // 开启服务
	Log(string)         // 日志输出
}

type Logger struct {

}
// 实现Service的Log方法
func (l *Logger) Log(s string)  {
	fmt.Println(s)
}
// 游戏服务
type GameService struct {
	Logger
}
// 实现Service的Start方法
func (g *GameService) Start()  {

}

func main()  {
	var s Service = new(GameService)
	s.Start()
	s.Log("hello")
}

