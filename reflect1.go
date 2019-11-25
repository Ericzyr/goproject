package main

import (
	"fmt"
	"reflect"
)


type Pc struct{
	Name string
	Age int
}

func (p Pc) Connect() {
	fmt.Println("hello world")
}




func main(){
	p := Pc{"apple",12}
	Info(p)

}



func Info(o interface{})  {
	//反射的类型
	t := reflect.TypeOf(o)
	fmt.Println(t)
	fmt.Println(t.Kind())
	if k:= t.Kind();k != reflect.Struct{
		fmt.Println("xx")
		return
	}

	v := reflect.ValueOf(o)

	//反射的值的类型
	for i:=0;i<t.NumField();i++{
		f:=t.Field(i)
		val:= v.Field(i).Interface()
		fmt.Printf("%6s : %v = %v\n",f.Name,f.Type,val)
	}
	//反射的方法的类型
	for r :=0;r<t.NumMethod();r++{
		m:= t.Method(r)
		fmt.Printf("%6s : %v",m.Name,m.Type)
	}
}
