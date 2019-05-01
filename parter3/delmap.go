package main

import "fmt"

func main()  {
	scene := make(map[string]int)
	scene["route"] = 66
	scene["alison"] = 100
	scene["china"] = 11

	delete(scene,"china")
	for k,v := range scene{
		fmt.Println(k,v)
	}
}
