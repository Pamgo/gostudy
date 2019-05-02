package main

import "sync"

var (
	// 一个演示用的映射
	valueByKey = make(map[string]int)
	// 保证使用映射时的并发安全的互斥锁
	valueByKeyGuard  sync.Mutex
)

func readValue(key string) int {
	valueByKeyGuard.Lock()

	v := valueByKey[key]

	valueByKeyGuard.Unlock()
	return v
}

// 使用defer进行简化
func readValule1(key string) int {

	valueByKeyGuard.Lock()

	// defer 后面的语句不会马上执行，而是延迟到函数结束时调用
	defer valueByKeyGuard.Unlock()
	return valueByKey[key]
}
