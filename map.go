package main

import (
	"fmt"
	"reflect"
)

var vs = []int {1,67,56,9}

var mapstrut=  map[int]int{1:21,2:8,3:17}

func main()  {

	fmt.Println(reflect.TypeOf(mapstrut),"map结构")
	fmt.Println(reflect.TypeOf(vs),"切片结构")

	for _,args := range vs{
		fmt.Println(args)
	}

	for k,v :=range mapstrut {
		fmt.Println(k,":",v)
		fmt.Println(reflect.TypeOf(mapstrut))
	}
}
