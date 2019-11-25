package main

import (
	"fmt"
)

func main() {
	//runtime.GOMAXPROCS(runtime.NumCPU())
	//s:=runtime.GOMAXPROCS(runtime.NumCPU())
	//fmt.Println(s)
	c := make(chan bool)
	for i:=0;i < 10;i++ {
		go sumAnd(c,i)
	}
	for i:=0;i<10 ;i++  {
		<- c
	}

}


func sumAnd(c chan bool,index int)  {
	var sumy int
	for i:=0;i<1000 ;i++{
		sumy =sumy +i
	}
	fmt.Println(index, sumy)
	c <- true
}

