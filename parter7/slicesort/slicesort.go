package main

import (
	"sort"
	"fmt"
)

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

func main()  {
	heros := []*Hero{
		&Hero{"吕布",Tank},
		&Hero{"礼拜", Assassin},
		&Hero{"妲己", Mage},
		&Hero{"貂蝉", Assassin},
		&Hero{"关羽", Tank},
		&Hero{"诸葛亮", Mage},
	}
	sort.Slice(heros, func(i, j int) bool {
		if heros[i].Kind != heros[j].Kind {
			return heros[i].Kind < heros[j].Kind
		}
		return heros[i].Name < heros[j].Name
	})

	for _,v := range heros{
		fmt.Printf("%+v\n",v)
	}
}