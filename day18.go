package main

import (
	"fmt"
	"time"
)

var w string

var arry = []int{1,2,8,9,12}

func main()  {
	for _,a:=range arry{
		fmt.Print(a )
	}

	w="hwqrf"

	fmt.Println("this weekday ,sed APT test report every day ")
	fmt.Println(time.Now())
	fmt.Println(w)
}
