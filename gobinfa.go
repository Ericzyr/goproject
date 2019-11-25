package main

import (
	"fmt"
	"runtime"
)

var Pr =make(chan bool)

func main(){
	runtime.GOMAXPROCS(runtime.NumCPU())
	for k:=0;k<10 ;k++  {
		go sp(k)
	}
	for i:=0;i<10 ;i++  {
		<- Pr
	}
}
func sp(index int) {
	a := 1
	for i := 0; i <10; i++ {
		a += i
	}
	fmt.Println(index, a)
	Pr <- true
}