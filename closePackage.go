package main
import "fmt"

var arr=[]int {1,2,6,9}

type tea struct {
	studey string
}
type persiont struct{
	tea
	name string
	age  int
}

func ts(){
	at:= persiont{tea:tea{"adfa"},name:"asdf",age:12}
	at.name="xiaoming"
	at.age= 12

	fmt.Println(at.name,at.tea.studey)

}


func main(){
	st := Sum(5)
	st(2)
	fmt.Println(st(19))

}

//讲述函数中闭包，匿名函数的调用

func Sum(x int) func(y int) int {
	return func(y int) int {
		return x + y
	}

}

