package main

import (
	"fmt"
	"reflect"
)

var arryType []int

var numbers = make([]int,3)

var mapType map[int]string

func main()  {

	numbers = []int{2,1,8,4,7,725}
	fmt.Println(numbers,reflect.TypeOf(numbers),cap(numbers))

	arryType =[]int{1,3,12}
	fmt.Printf("%T\n",arryType)
	fmt.Println(cap(arryType))

	mapType = map[int]string{1:"jack",2:"tome"}
	fmt.Printf("%T\n",mapType)
}
