package main

import (
	"os"
	"os/exec"
)

func main()  {
	//c :=exec.Command("ping","www.baidu.com")
	c :=exec.Command("ipconfig")
	//
	c.Stdout = os.Stdout
	c.Run()
}