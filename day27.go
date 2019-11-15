package main

import "fmt"

type got struct {
	name string
	age int
}

func  athome(x int )string {
	a := got{}
	a.name = "asdf"
	a.age = 12

	return "sasdf"
}

func main() {
	fmt.Println(athome(5))
}