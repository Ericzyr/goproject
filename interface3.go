package main

import (
	"errors"
	"fmt"
)


func main(){
	fmt.Println(Sgt("ok").Name())


	fmt.Println(errors.New("sfasdf"))


	var u USB = &applephoto{"side"}
	fmt.Println(u.Name())
}

type USB interface {
	Name() string
}

func Sgt(text string)USB{
	return &applephoto{text}
}


type applephoto struct {
	s string
}

func (e *applephoto)Name()string {
	return e.s
}
