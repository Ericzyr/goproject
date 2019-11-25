package main

import (
	"fmt"
	"reflect"
)


type Pc struct{
	Name string
	Age int
}
type manager struct{
	Pc
	address string
}


func main(){
	p := manager{Pc:Pc{"jack",18},address:"beijing"}
	t :=reflect.ValueOf(p)
	//匿名函数的调用
	//fmt.Println(t.FieldByIndex([]int{0,1}))
	fmt.Println(t.Field(1))

}
