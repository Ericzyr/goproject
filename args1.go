package main

import (
	"AddTest"
	"fmt"
)

func main(){
	run("fruit","apple")

	fmt.Println(AddTest.Total())
}

func run(param ...string)  {
	if len(param)>0{
		for i:=0;i<len(param);i++ {
			fmt.Println(param[i])
		}
	} else{
		fmt.Println("no get param")
	}
}