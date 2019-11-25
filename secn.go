package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func main()  {
	v:=runtime.GOOS
	fmt.Println(v)
	//c :=exec.Command("ping","www.baidu.com")
	c :=exec.Command("ifconfig")
	c.Stdout = os.Stdout
	c.Run()
}