package main

import "fmt"

func main()  {
	A()
	B(0)
	C()
}

func A()  {
	fmt.Println("func A")
}

func B(args int8)  {

	defer func() {
		fmt.Println(recover())
		if err:= recover();err!=nil{
			fmt.Println(err)
		}
	}()
	fmt.Println(10/args)
}

func C()  {
	fmt.Println("func C")
}