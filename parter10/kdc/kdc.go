package main

import (
	"fmt"
	"reflect"
)

func main()  {
	type dog struct {
		LegCount int // 大写首字母
	}

	valueOfDog := reflect.ValueOf(&dog{})

	valueOfDog = valueOfDog.Elem();

	vLegCount := valueOfDog.FieldByName("LegCount")
	vLegCount.SetInt(10)

	fmt.Println(vLegCount.Int())
}
