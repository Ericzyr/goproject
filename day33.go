package main

import "fmt"

func vegetables(x int )func(y int )int {
	return func(y int)int{
		return x + y
	}
}

type golanguage struct {
	name string
	age int
	address string
}

func (c *golanguage) study(){
	c.name= "cat"
	c.age=12
	fmt.Println(c.name,c.age)
}


func main() {
	st :=vegetables(2)(5)
	fmt.Println(st)

	c := golanguage{}
	c.study()


}
