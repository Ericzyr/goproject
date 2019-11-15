package main

import "fmt"

type computer struct {
	tools string
}

func (a computer) OpenWeb(fileter string ){
	fmt.Print("use ",a.tools," watch ",fileter)

}

func main()  {
	c1 := computer{}
	c1.tools = "keyboard"
	c1.OpenWeb("baidu")

}
