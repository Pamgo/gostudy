package main

import (
	"fmt"
)

type Wheel struct {
	Size int
}

type Car struct {
	Wheel
	Engine struct{
		Power int
		Type string
	}
}

func main()  {
	c := Car{
		Wheel:Wheel{
			Size:18,
		},
		Engine: struct {
			Power int
			Type  string
		}{Power: 143, Type: "1.4t"},
	}
	fmt.Printf("%+v\n",c)
}

