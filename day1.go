package main

import (
	"fmt"
	"os"
	"time"
)
//常量的定义
const PI  = 3.14

//一般类型的声明
type per string

//全局变量的声明与赋值
var agee int = 12
var golangust string="English"

//结构的声明
type persion struct {
	name string
	age int
}
var P persion

//接口的声明
type golang interface {}

func main()  {
	a :=persion{name:"jack",age:15}
	a.name="jack"
	P.name="josh"
	fmt.Println(a.age)
	fmt.Println(per("wqfdsaf"))
	fmt.Println("hello world")
	fmt.Println(os.Hostname())
	fmt.Println(time.Now().Date())
	forgo()
}

func forgo()  {
	for i :=0;i<10;i++ {
		fmt.Print(i)
		if i == 5{
			break
		}
	}
	a := 5
	b := 7
	if a > b{
		fmt.Println("比较大小",a,">",b)
	} else {
		fmt.Println("比较大小",b,"<",a)
	}

}