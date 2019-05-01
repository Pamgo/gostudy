package main

import "fmt"

// 函数多返回值
func typedTwoValues() (int,int)  {
	return 1,2
}

func main()  {
	a,b := typedTwoValues()
	fmt.Println(a,b)
}
