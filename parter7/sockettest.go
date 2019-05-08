package main

import "io"

type Socket struct {
	
}

func (s *Socket) Write(p []byte) (n int, err error)  {
	return 0,nil
}

func (s *Socket) Close() error {
	return nil
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}
// 使用io.writer的代码，并不知道Socket和io.Closer的存在
func usingWriter(writer io.Writer)  {
	writer.Write(nil)
}
// 使用io.Closer，并不知道Socket和io.writer的存在
func usingCloser(closer io.Closer)  {
	closer.Close()
}

func main()  {
	s := new(Socket)

	usingWriter(s)

	usingCloser(s)
}