package main

import (
	"bufio"
	"fmt"
	"os"
)

const ROUTE float32 = 3.14

func main()  {
	fmt.Print("输入圆的半径")
	input:=bufio.NewScanner(os.Stdin)
	input.Scan()
	r:=input.Text()
	sumt := ROUTE*r*r

}