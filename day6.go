package main

import "fmt"

func main() {
		a,b := sum(9,2)
		fmt.Println("two number plus equal to ",a)
		fmt.Println("two number minus equal to ", b)
}

func sum(a int, b int)(int,int ){
	plus := a + b
	minus:= a - b
	return plus,minus
}

//匿名函数
/*
func main()  {
	a :=func(x int)int{
		return x
	}
	fmt.Println(a(9))
}
*/