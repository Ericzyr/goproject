package main

import "fmt"

//func s() int64{
//	fmt.Println("function s() --->")
//	return 1
//}
//var kv int64 =s()

var precomputed=[5]int{}
//init 初始化数据
func init(){
	//var current int=1
	precomputed[0]=2
	//for i:=1;i<len(precomputed);i++{
	//	precomputed[i]=precomputed[i-1]*1
	//}
}

func main(){
	fmt.Println(precomputed)
}

//var v int

//func init()  {
//	for i:=0;i<10;i++{
//		v =v + i
//	}
//}
//

//
//func main ()  {
//	fmt.Print(v)
//}