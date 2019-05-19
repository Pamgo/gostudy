package main

import (
	"bufio"
	"net"
)

// 连接回话逻辑
func handleSession(conn net.Conn, call func(net.Conn,[]byte) bool)  {

	// 创建一个sokcet读取器
	dataReader := bufio.NewReader(conn)

	for  {
		// 从连接读取封包
		pkt, err := readPacket(dataReader)

		// 从回调外部
		if err != nil || !call(conn, pkt.Body) {
			// 回调要求退出
			conn.Close()
			break
		}
	}
}
