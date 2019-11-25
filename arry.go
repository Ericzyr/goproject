package main

import (
	"fmt"
	"os"
)

func main()  {
	var n = []int {10,6,89,34,28}
	for i:=0;i<len(n);i++{
		fmt.Println(n[i])
	}

	for _,v:=range n{
		fmt.Print(v , "、")
	}
	t,erro:=os.Getwd()
	if erro !=nil {
		fmt.Println(erro)
	}
	fmt.Println(t)
}
