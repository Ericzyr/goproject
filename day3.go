package main

import "fmt"

func main()  {
	a:=2
	switch{
	case a > 0:
		fmt.Println("a=1")
		fallthrough
	case a >= 2:
		fmt.Println("a=2")
	default:
		fmt.Println("none")
	}
	loop()
}

func loop()int {
	var a  = 3
	return a
}