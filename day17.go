package main

import (
	"fmt"
)

type Econ struct {
	name string
	age int
}


func (a Econ) meche() int {
	fmt.Println(a.name,"hello world")
	return 2
}

func t() int  {
	return 123
}

func main()  {
	as := Econ{name:"jcke",age:19}
	fmt.Println(as)
	as.age=12
	as.name="tom"
	as.meche()
	fmt.Println(t())
}