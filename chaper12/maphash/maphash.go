package main

import (
	"fmt"
)

type Profile struct {
	Name string
	Age int
	Married bool
}

func SimpleHash(str string) (ret int) {

	for i := 0; i< len(str); i++ {
		c := str[i]
		ret += int(c)
	}
	return
}

// 查询键
type ClassicQueryKey struct {
	Name string
	Age int
}
// 计算查询键的哈希值
func (c *ClassicQueryKey) hash() int {
	return SimpleHash(c.Name) + c.Age * 1000000
}
// 创建哈希值到数据的索引关系中
var mapper = make(map[int][]*Profile)

// 构建数据索引
func buildIndex(list []*Profile)  {
	for _,profile := range list {
		key := ClassicQueryKey{profile.Name,profile.Age}

		existValue := mapper[key.hash()]

		existValue = append(existValue, profile)

		mapper[key.hash()] = existValue
	}
}

type queryKey struct {
	Name string
	Age int
}
var mapper1 = make(map[queryKey]*Profile)
func buildIndex1(list []*Profile)  {
	for _, profile := range list{
		key := queryKey{
			Name:profile.Name,
			Age:profile.Age,
		}

		mapper1[key] = profile
	}
}
func queryData(name string, age int)  {

	keyToQuery := ClassicQueryKey{name,age}

	resultList := mapper[keyToQuery.hash()]

	for _, result := range resultList {
		if result.Name == name && result.Age == age {
			fmt.Println(result)
			return
		}
	}

	fmt.Println("no found")
}

func main()  {
	list := []*Profile{
		{Name:"alison",Age:30,Married:true},
		{Name:"alison2",Age:20},
		{Name:"alison3",Age:21,Married:false},
	}

	buildIndex(list)

	queryData("alison2",20)
}
