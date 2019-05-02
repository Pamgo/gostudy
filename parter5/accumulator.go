package main

import "fmt"

func Accumulate(value int) func() int  {
	return func() int {
		value++
		return value
	}
}

func main()  {
	acc := Accumulate(1)

	fmt.Println(acc())
	fmt.Println(acc())

	acc2 := Accumulate(10)
	fmt.Println(acc2())
}
