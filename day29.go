package main

import (
	"fmt"
	"os/exec"
)

func main()  {

	a:=exec.Command("pwd")
	fmt.Println(*a)

}
