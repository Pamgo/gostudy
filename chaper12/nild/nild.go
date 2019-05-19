package main

import "fmt"

type MyImplement struct {

}

func (m *MyImplement) string() string {
	return "hi"
}

func GetStringer() fmt.Stringer {
	var s *MyImplement = nil

	return s
}

func main()  {
	if GetStringer() == nil {
		fmt.Println("Getstringer == nil")
	} else {
		fmt.Println("getStringer != nil")
	}
}
