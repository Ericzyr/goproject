package main

import "fmt"

type human struct {
	name string
	age int
}

type teacher struct{
	human
	teacherBook string
}

type student struct {
	human
	book string
}

func main ()  {
	s := student{human:human{"ad",12},
		book:"语文book",
	}
	fmt.Println(s.book)

}