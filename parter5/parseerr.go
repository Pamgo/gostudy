package main

import "fmt"

type ParseError struct {
	Filename string   // 文件名
	Line     int      // 行号
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("%s:%d",e.Filename,e.Line)
}

func newParseError(filename string, line int) error  {
	return &ParseError{filename,line}
}

func main()  {
	var e error

	// 创建一个错误实例，保护文件名和行号
	e = newParseError("main.go",1)

	fmt.Println(e.Error())

	switch detail := e.(type) {
	case *ParseError:
		fmt.Printf("Filename: %s Line:%d\n", detail.Filename,detail.Line)
	default:
		fmt.Println("othere error")
	}
}



