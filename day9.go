package main

import (
	"fmt"
)

type A struct{
	name string
	age int
}
func (a A) Method(d int ){
	a.name= "Tom"
	fmt.Println(a.name,"A")
}

func main()  {
	c:= A{
		name:"jack",
		age : 12,
	}
	//1
	c.Method(8)
	//2
	fmt.Println(c.name)
}

