package main

import (
	"fmt"
	"sync"
)

var (
	// 逻辑中使用的某个变量
	count int
	// 与变量对应的使用互斥锁
	//countGuard sync.Mutex
	countGuard sync.RWMutex
)

func GetCount() int {
	countGuard.Lock()

	defer countGuard.Unlock()

	return count
}

func SetCount(c int)  {
	countGuard.Lock()
	count = c
	countGuard.Unlock()
}

func GetCount2() int {
	countGuard.Lock()
	defer countGuard.Unlock()

	return count
}

func main()  {
	SetCount(1)

	fmt.Println(GetCount2())
}
