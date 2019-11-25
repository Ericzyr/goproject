package  main

import (
	"fmt"
)

type A struct{
	name string
}
type B struct {
	name string
}

func (a *A)Print(){
	a.name="AA"
	fmt.Println("A")

}
func(b B) Print(){
	b.name="BB"
	fmt.Println("B")
}

func main()  {
	a := new(A)
	a.Print()
	fmt.Println(a.name)

	b:= new(B)
	b.Print()
	fmt.Println(b.name)

}