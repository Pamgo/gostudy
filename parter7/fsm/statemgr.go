package main

import "errors"

// 状态没有找到错误
var ErrStateNotFound = errors.New("state not found")
// 禁止在同状态间转移
var ErrForbidSameStateTransit = errors.New("forbid same state transit")
// 不能转移到指定状态
var ErrCannotTransitToState = errors.New("cannot transit to state")

// 状态管理器
type StateManager struct {
	// 已经添加的状态
	stateByName map[string]State

	// 状态改变时的回调
	OnChange func(from, to State)

	// 当前状态
	curr State
}

func (sm *StateManager) Get(name string) State {
	if v, ok := sm.stateByName[name]; ok {
		return v
	}
	return nil
}
// 添加一个状态到管理器中
func (sm *StateManager) Add(s State) {

	// 获取状态的名称
	name := StateName(s)

	// 将s转换为能设置名称的接口，然后调用该接口
	s.(interface{
		setName(name string)
	}).setName(name)

	// 根据状态名获取已经添加的状态，检查该状态是否存在
	if sm.Get(name) != nil {
		panic("duplicate state:"+name)
	}
	// 根据名字保存到map中
	sm.stateByName[name] = s
}

func NewStateManager() *StateManager {
	return &StateManager{
		stateByName:make(map[string]State),
	}
}

/***********在状态间转移***************/

func (sm *StateManager) CurrState() State {
	return sm.curr
}

func (sm *StateManager) CanCurrTransitTo(name string) bool {
	if sm.curr == nil {
		return true
	}

	// 相同状态下转换
	if sm.curr.Name() == name && !sm.curr.EnableSameTransit() {
		return false
	}
	// 使用当前状态，检查是否转移到指定名字状态
	return sm.curr.CanTransitTo(name)
}

// 转移到指定状态
func (sm *StateManager) Transit(name string) error {
	// 获取目标状态
	next := sm.Get(name)

	// 目标不存在
	if next == nil {
		return ErrStateNotFound
	}

	// 记录转移前的状态
	pre := sm.curr

	// 当前有状态
	if sm.curr != nil {
		// 相同的状态不用转换
		if sm.curr.Name() == name && !sm.curr.EnableSameTransit() {
			return ErrForbidSameStateTransit
		}

		// 不能转移到目标状态
		if !sm.curr.CanTransitTo(name) {
			return ErrCannotTransitToState
		}
		sm.curr.OnEnd()
	}

	// 将当前状态切换为要转移的目标状态
	sm.curr = next

	// 调用新状态开始
	sm.curr.OnBegin()

	if sm.OnChange != nil {
		sm.OnChange(pre, sm.curr)
	}

	return nil
}