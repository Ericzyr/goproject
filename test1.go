package main

import (
	"fmt"
	"reflect"
	"regexp"
)

var regular = "Hell 世界！123 Go."


func main()  {
	// 查找连续的小写字母
	fmt.Println(reflect.TypeOf(regular))

	reg := regexp.MustCompile("[a-z]+")
	w:=reg.FindAllString(regular, -1)
	fmt.Println(w)

	// 查找行首以 H 开头，以空白结尾的字符串（非贪婪模式）
	reg = regexp.MustCompile("(?U)l.*o")
	fmt.Printf("%q\n", reg.FindAllString(regular, -1))
	// ["Hello" "Go"]

}