package  main

import "fmt"

type good struct {
	fruit string
	vegetable string
	number int
}

func (g *good) todoing() {
	fmt.Print("this is doing lift")
}


func main() {
	c := new(good)
	c.todoing()
}
