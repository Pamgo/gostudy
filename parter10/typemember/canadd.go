package main

import (
	"fmt"
	"reflect"
)

func main()  {
	var a int = 1024

	valueOfA := reflect.ValueOf(&a)

	valueOfA = valueOfA.Elem()

	valueOfA.SetInt(1)

	fmt.Println(valueOfA.Int())
}
