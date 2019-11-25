package main

import (
	"fmt"
	"os/exec"
)

func main()  {
	a:=exec.Command("ifconfig")
	fmt.Println(*a)

}
