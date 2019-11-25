package main

import (
	"errors"
	"fmt"
)

func tst(i int)(int,error){

	if i>0{
		return i,errors.New("i>0")
	}
	return i,errors.New("")
}

func main()  {
	n,v:=tst(0)
	fmt.Print(n,v)
}