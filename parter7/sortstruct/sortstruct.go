package main

import (
	"sort"
	"fmt"
)

// 声明英雄的分类
type HeroKind int

// 定义HeroKind常量，类似于枚举
const (
	None HeroKind = iota
	Tank
	Assassin
	Mage
)
// 定义英雄名单的结构
type Hero struct {
	Name string
	Kind HeroKind
}

// 将英雄的指针的切片定义为hero类型
type Heros []*Hero

// 实现sort.interface的取元素的方法
func (h Heros) Len() int {
	return len(h)
}

// 实现sort.Interface接口比较元素方法
func (h Heros) Less(i, j int) bool {
	// 如果英雄分类不一致时，优先对分类进行排序
	if h[i].Kind != h[j].Kind {
		return h[i].Kind < h[j].Kind
	}

	return h[i].Name < h[j].Name
}

// 实现sort.interface接口交换元素的方法
func (h Heros) Swap(i,j int)  {
	h[i],h[j] = h[j],h[i]
}

func main()  {
	heros := Heros{
		&Hero{"吕布",Tank},
		&Hero{"礼拜", Assassin},
		&Hero{"妲己", Mage},
		&Hero{"貂蝉", Assassin},
		&Hero{"关羽", Tank},
		&Hero{"诸葛亮", Mage},
	}
	sort.Sort(heros)

	for _,v := range heros{
		fmt.Printf("%+v\n",v)
	}
}