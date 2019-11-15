package main

import "fmt"

type person struct{
	Name string
	Age int
	Contact struct{
		Phone,City string
	}
}

func main()  {
	a := person	{}
	a.Name="jack"
	a.Age = 12
	a.Contact.Phone = "1234345345"
	fmt.Println(a)
}