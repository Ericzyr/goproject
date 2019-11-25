package main

import "fmt"

type Readbook interface{
	read()
}
type English struct{
	name string
	age int
	adress string
}

type History struct {}

func (a English) read(){
	a.name="jack"
	a.age =12
	a.adress="beijing"
	fmt.Println(a.name,"Heolo to ",a.adress )
}

func (b History)read(){

}

var arrtw =[]int{2,4,9}
var aryy2  =[]int{2,9,6}

var m = map [int]interface{}{1:"sfd",2:"sae"}


func main(){
	var book Readbook
	book = English{}
	book.read()

	var book2 Readbook
	book2 = History{}
	book2.read()

	for _,v:=range m{
		vt:=v
		fmt.Println(vt)
	}

}