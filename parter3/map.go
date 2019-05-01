package main

import "fmt"

func main()  {
	scene := make(map[string]int)
	scene["route"] = 66
	fmt.Println(scene["route"])
	v := scene["route2"]
	fmt.Println(v)
	v, ok := scene["route"]
	if ok {
		fmt.Println(v,ok)
	}
	// 遍历
	scene["route"] = 77
	scene["alison"] = 3
	scene["china"] = 960

	for k,v := range scene  {
		fmt.Println(k,v)
	}

}
