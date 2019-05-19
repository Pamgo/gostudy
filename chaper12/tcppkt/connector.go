package main

import (
	"fmt"
	"net"
	"strconv"
)

// 连接器，传入连接地址和发送封包次数
func Connector(address string, sendTimes int)  {

	// 尝试用Socket连接地址
	conn, err := net.Dial("tcp",address)

	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < sendTimes; i++ {
		str := strconv.Itoa(i)  // Itoa函数可以将数值转为字符串

		if err := writePacket(conn, []byte(str)); err != nil {
			fmt.Println(err)

			return
		}
	}
}
