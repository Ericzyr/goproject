package main

import (
	"fmt"
	"reflect"
	"strconv"
)

//var a float64
//func main() {
//	a = 123.4
//	fmt.Println("a type is ", reflect.TypeOf(a))  //打印 a type is float64
//	b := int(a)
//	fmt.Println(b)                                //打印 123
//	fmt.Println("b type is ", reflect.TypeOf(b))  //打印 b type is uint32
//}

//func main() {
//	a := '中'
//	fmt.Println(reflect.TypeOf(a),&a,a)
//
//	b := byte(a)
//	fmt.Println(reflect.TypeOf(b),&b,b)
//}


//func main() {
//	//a := "67"
//	//b, err := strconv.Atoi(a)
//	//if err != nil {
//	//	fmt.Println(err)
//	//}
//	//fmt.Println(reflect.TypeOf(b),b)
//	//
//	s := "elevator"
//
//	p := &s
//	fmt.Println(p)

	//c:= strconv.Itoa(*p)
	//fmt.Println(reflect.TypeOf(c),c)

	//var p int64 = 'a'
	//string1 := strconv.FormatInt(p,16)
	//fmt.Println(string1)
	//
	//string2 := strconv.Itoa(d)
	//fmt.Println(string2)

//}




//func stringToBytes (s string) []byte {
//	return []byte(s)
//}
//
//
//func main() {
//	s := "abc"
//	bytes := stringToBytes(s)
//	fmt.Println(bytes)
//	// 打印 [97 98 99]
//}


func stringToBytes (s string) []byte {
	return []byte(s)
}

func main() {
	var p int64 = 'a'
	string1 := strconv.FormatInt(p,10)
	fmt.Println(string1)

	a := 'a'
	fmt.Println(reflect.TypeOf(a),a)
	b := string(97)
	fmt.Println(reflect.TypeOf(b),b)

	s := "shl"
	var	arryA = [] string{}
	bytes := stringToBytes(s)
	for _,v :=range  bytes{
		arryA=append(arryA, string(v))
	}
	fmt.Println(arryA)
	fmt.Printf("%T\n",arryA)
}