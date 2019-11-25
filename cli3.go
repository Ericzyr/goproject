package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("input user name: ")
	input := bufio.NewScanner(os.Stdin)//初始化一个扫表对象
	input.Scan()
	line := input.Text()//把输入内容转换为字符串

	fmt.Println("input user pwd: ")
	input2 := bufio.NewScanner(os.Stdin)//初始化一个扫表对象
	input2.Scan()
	line2 := input2.Text()//把输入内容转换为字符串

	fmt.Println(line,line2)//输出到标准输出





	//for input.Scan() {//扫描输入内容
	//	line := input.Text()//把输入内容转换为字符串
	//	fmt.Println(line)//输出到标准输出
	//}
}
