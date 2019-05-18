package main

import (
	"fmt"
	"reflect"
)

type dummy struct {
	a int
	b string

	float32
	bool
	next *dummy
}

func main()  {
	d := reflect.ValueOf(dummy{
		next:&dummy{},
	})

	fmt.Println("NumField", d.NumField())

	floatField := d.Field(2)

	fmt.Println("Field", floatField.Type())

	fmt.Println("FieldByName:", d.FieldByName("b").Type())

	fmt.Println("fieldbyIndex:", d.FieldByIndex([]int{4,0}).Type())



}
