package main

import (
	"fmt"
	"net"
	"sync"
)

// 接收器
type Acceptor struct {

	// 保存侦听器
	l net.Listener

	// 侦听器的停止同步
	wg sync.WaitGroup

	// 连接数据回调
	OnSessionData func(net.Conn, []byte) bool
}

func (a *Acceptor) listen(address string)  {

	// 侦听开始，添加一个任务
	a.wg.Add(1)

	// 在退出函数时，结束侦听任务
	defer a.wg.Done()

	var err error

	// 根据给定地址进行侦听
	a.l, err = net.Listen("tcp", address)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// 侦听循环
	for {
		// 新连接还没到来时，Accept是阻塞的
		conn, err := a.l.Accept()

		if err != nil {
			break
		}

		// 根据连接开启回话，这个过程需要并行执行
		go handleSession(conn, a.OnSessionData)

	}
}

// 异步开始侦听
func (a *Acceptor) Start(address string)  {
	go a.listen(address)
}

func (a *Acceptor) Stop() {
	a.l.Close()
}

func (a *Acceptor) Wait() {
	a.wg.Wait()
}

// 实例化一个侦听器
func NewAcceptor() *Acceptor {
	return &Acceptor{}
}