package main

import (
	"bytes"
	"reflect"
)

// 将切片转换为JSON格式并输出到缓冲中
func writeSlice(buff *bytes.Buffer, value reflect.Value) error {

	// 写入切片开始标记
	buff.WriteString("[")
	// 遍历每一个切片元素
	for s := 0; s < value.Len() ; s++ {
		sliceValue := value.Index(s)

		// 写入每一个切片元素
		writeAny(buff, sliceValue)
		// 每一个元素尾部写入逗号，最后一个字段不添加
		if s < value.Len()-1 {
			buff.WriteString(",")
		}
	}
	buff.WriteString("]")
	return nil
}
