package main

import (
	"fmt"
)

type Phone interface {
	call()
}

type IPhone struct {
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}


func main() {
	var phone Phone

	phone = NokiaPhone{}
	phone.call()

	phone = IPhone{}
	phone.call()
}
