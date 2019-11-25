package main

import (
	"fmt"
	"os"
	"strings"
)

func main()  {
	v,err := os.Getwd()
	if err!=nil {
		fmt.Println("get path err",err)
	}
	fmt.Println(v+"\\a.txt")

	var b string  = "adfsadfdsafasf"

	if strings.Contains(b ,"s.") == true{
		fmt.Println("包括")
	}else {
		fmt.Println("other")
	}


}


