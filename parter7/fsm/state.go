package main

import "reflect"

// 状态接口
type State interface {

	Name() string  // 获取状态名称

	EnableSameTransit() bool // 该状态是否允许同状态转移

	OnBegin()   // 响应状态开始

	OnEnd() // 响应状态结束时

	CanTransitTo(name string) bool // 判断是否能转移某个状态
}

// 从状态实例获取状态名
func StateName(s State) string {
	if s == nil {
		return "none"
	}

	// 使用反射获取状态名称
	return reflect.TypeOf(s).Elem().Name()
}
