package main

import (
	"fmt"
	"reflect"
	"strings"
)

var str = []string{"hello", "world", "company","is","sdf"}

var str2 string = "hello world  or ,  two"
func main()  {
	//:Join 将 a 中的子串连接成一个单独的字符串，子串之间用 sep 分隔
	arr:= strings.Join(str,",")
	fmt.Println(arr)
	//字符串替换
	tp:=strings.Replace(arr,","," ",-1)
	fmt.Println(tp)

//	==============================================
//将字符串 s 以空格为分隔符拆分成若干个字符串，若成功则返回分割后的字符串切片
	arr2:=strings.Fields(str2)
	fmt.Println(reflect.TypeOf(arr2),arr2)


	var arr3 []string

	for i:=0;i<len(arr2);i++ {
		if strings.EqualFold(arr2[i],","){
			break
		}
		arr3=append(arr3,arr2[i])
	}
	fmt.Println(arr3)
}