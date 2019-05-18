package main

import (
	"encoding/json"
	"fmt"
)

func main()  {
	// 声明技能结构体
	type Skill struct {
		Name string
		Level int
	}

	// 声明角色结构
	type Actor struct {
		Name string
		Age int

		Skills []Skill
	}

	a := Actor{
		Name:"alison",
		Age:26,
		Skills:[]Skill{
			{Name:"pamgo", Level:1},
			{Name:"byteRun",Level:2},
			{Name:"byte",Level:3},
		},
	}

	//if result, err := MarshalJson(a); err == nil {
	//	fmt.Println(result)
	//} else {
	//	fmt.Println(err)
	//}

	if result1, err := json.Marshal(a); err == nil {
		fmt.Println(string(result1))
	}
}
