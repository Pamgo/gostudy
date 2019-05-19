package main

import (
	"bytes"
	"encoding/binary"
	"io"
)

type Packet struct {
	Size uint16    // 包体大小
	Body []byte    // 包体数据
}
// 将数据写入dataWriter
func writePacket(dataWriter io.Writer, data []byte) error {
	// 准备一个字节数组缓存
	var buffer bytes.Buffer

	// 将Size写入缓存
	if err := binary.Write(&buffer, binary.LittleEndian, uint16(len(data))); err != nil {
		return err
	}

	// 写入包体数据
	if _, err := buffer.Write(data); err != nil{
		return err
	}

	// 获得写入的完整数据
	out := buffer.Bytes()

	// 写入dataWriter
	if _, err := dataWriter.Write(out); err != nil {
		return err
	}

	return nil
}
// 从dataWriter中读取包
func readPacket(dataReader io.Reader) (pkt Packet, err error)  {

	// Size为uint16类型，占2个字节
	var sizeBuff = make([]byte, 2)

	// 持续读取Size直到读到为止
	_,err = io.ReadFull(dataReader, sizeBuff)

	if err != nil {
		return
	}

	// 使用bytes.Reader读取sizeBuffer中的数据
	sizeReader := bytes.NewReader(sizeBuff)

	// 读取小端uint16作为size
	err = binary.Read(sizeReader, binary.LittleEndian, &pkt.Size)

	if err != nil {
		return
	}
	// 分配包体大小
	pkt.Body = make([]byte, pkt.Size)

	_, err = io.ReadFull(dataReader, pkt.Body)
	return
}