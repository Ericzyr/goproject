package main

import (
	"fmt"
	"reflect"
)

const IP = 3.14159262324224
var st int

func main()  {
	var a float64= IP*5
	fmt.Println(a,reflect.TypeOf(a))
}
