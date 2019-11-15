package main

import (
	"fmt"
)

const API = "api.edgescale.com"

var mapw = map[string]int{"a":8,"b":9,"c":10,"d":100}


var arryw = []int{2,5,2,8,9}

func main() {
	for k,v:=range mapw{
		fmt.Println(k,v)
	}
}