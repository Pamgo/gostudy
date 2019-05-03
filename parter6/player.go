package player

import ""

type Player struct {
	currPos Vec2   // 当前位置
	targetPos Vec2    // 目标位置
	speed float32 // 移动速度
}

// 设置玩家移动的目标位置
func (p *Player) MoveTo(v Vec2)  {
	p.targetPos = v
}

// 获取当前位置
func (p *Player) Pos() Vec2 {
	return p.currPos
}

// 判断是否到达目的地
func (p *Player) IsArrived() bool {
	return p.currPos.Distance(p.targetPos) < p.speed
}

// 更新玩家位置
func (p *Player) Update() {
	if !p.IsArrived() {
		// 计算出当前位置指向目标的朝向
		dir := p.targetPos.Sub(p.currPos).Normalize()

		// 添加速度矢量生成新的位置
		newPos := p.currPos.Add(dir.Scale(p.speed))

		p.currPos = newPos
	}
}
// 创建新玩家，构造器模式
func NewPlayer(speed float32) *Player {
	return &Player{speed:speed}
}
