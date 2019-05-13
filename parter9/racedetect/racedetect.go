package main

import (
	"fmt"
	"sync/atomic"
)

var (
	seq int64
)

func GenID() int64 {

	return atomic.AddInt64(&seq, 1)

	//return seq
}

func main()  {
	for i := 0; i < 10; i++ {
		go GenID()
	}

	fmt.Println(GenID())
}
