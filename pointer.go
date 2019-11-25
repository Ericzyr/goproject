package main

import "fmt"

func main()  {
	//定义变量
	var num int =10
	fmt.Println(num)

	//定义指针
	var ptr *int

	//pointer 赋值 取变量的地址&变量
	ptr = &num
	fmt.Println(&num)

	//取变量的地址的值，用指针取*指针
	fmt.Println(*ptr)

	//根据指针修改值
	*ptr=20
	fmt.Println(num)

}
