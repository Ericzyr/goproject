package main

import "fmt"

var a = []string{"a","b","c","d","e"}

func main()  {
	c := a[1:3]
	fmt.Println(c,len(c), cap(c))
}