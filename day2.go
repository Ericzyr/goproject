package main
import (
	"fmt"
)

var b int16
var c string


type perw struct {
	name string
	age  int64
}

//set GOARCH=amd64
//set GOOS=linux

func main()  {
	pv :=new(perw)
	pv.name="jack"
	pv.age=12
	ab := []int {2,3,56,23}
	for i:=0; i<len(ab);i++  {
		fmt.Print(ab[i],"\t")
	}
}
