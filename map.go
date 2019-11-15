package main

import (
	"fmt"
	"reflect"
)

var s = []int {1,67,56,9}


var mapstrut=  map[int]int{1:21,2:8,3:17}

func main()  {

	fmt.Println(reflect.TypeOf(mapstrut))
	for i:=0;i<len(s);i ++ {
		fmt.Println(s[i])
	}

	for k,v :=range mapstrut {
		fmt.Println(k,":",v)
		fmt.Println(reflect.TypeOf(mapstrut))
	}
}
