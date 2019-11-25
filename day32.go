package main

import "fmt"

var t int = 1000

func main()  {
	a:=&t
	t=18
	fmt.Println(*a)
}