package main

import "fmt"

type Bag struct {
	items []int
}

func (b *Bag) Insert(itemid int)  {
	b.items = append(b.items, itemid)
}

func main()  {
	b := new(Bag)
	b.Insert(1001)
	b.Insert(1002)

	for _,v := range b.items{
		fmt.Println(v)
	}

}
