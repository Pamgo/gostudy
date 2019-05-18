package main

import (
	"bytes"
	"reflect"
)

func MarshalJson(v interface{}) (string, error) {

	// 准备一个缓存
	var b bytes.Buffer

	if err := writeAny(&b, reflect.ValueOf(v)); err == nil {
		return b.String(),nil
	} else {
		return "", nil
	}
}
