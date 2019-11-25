package  main

import (
	"fmt"
)

type vegetable struct{
	name string
	age int
}
func (f *vegetable) drink()  {
	f.name="2banana"
	fmt.Println(f.name)
}
func main()  {
	a :=vegetable{}
	a.name="1apple"
	fmt.Println(a.name)
	a.drink()
	fmt.Print(a.name)
}