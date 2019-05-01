package main

import "fmt"

type Person struct {
	Age int
}

func (p *Person) GrowUp() {
	p.Age++
}

func main()  {
	m := map[string]*Person {
		"x" : &Person{
			Age:5,
		},
	}
	m["x"].Age = 23
	m["x"].GrowUp()
	fmt.Println(m["x"].Age)

/*	m := map[string]Person{
		"zhansan":Person{
			Age:5,
		},
	}

	m["zhansan"].Age=23*/

}


