package main

import (
	"fmt"
	"os"
	"runtime"
)

func main()  {

	var system =runtime.GOOS
	if system == "windows"{
		fmt.Println("system is Windows")
	} else if system == "linux" {
		fmt.Println("system is Linux")
	}

	host, err := os.Hostname()
	if err != nil {
		fmt.Printf("%s", err)
	} else {
		fmt.Printf("hostname: %s", host)
	}
}