package main

import "fmt"

type gofile struct {
	name string
	age int
}

func (g *gofile)look( nw string){
	fmt.Print("wqfdas "+ nw)
}
func main(){
	a:=gofile{"asd",12}
	a.look("go have")
}