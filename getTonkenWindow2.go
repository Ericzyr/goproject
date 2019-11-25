package main
import (
	"bufio"
	"encoding/json"
	"fmt"
	"HttpRequest"
	"log"
	"os"
)


var (
	useragent ="Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.110 Safari/537.36";
	ContentType ="application/json"
)

func Token( Url string,myname string ,mypwd string)string {
	req := HttpRequest.NewRequest()
	req.SetHeaders(map[string]string{
		"user-agent": useragent,
		"Content-Type": ContentType,
	})

	var postData = map[string]interface{}{
		"username": myname,
		"password": mypwd,
	}
	resp, err := req.Post(Url + "/v1/users/login",postData)

	if err != nil {
		log.Println(err)
		return ""
	}
	if resp.StatusCode() == 200 {
		body, err := resp.Body()
		if err != nil {
			log.Println(err)
			return ""
		}
		var user = map[string]string{}
		json.Unmarshal(body,&user)
		if user["message"] == "Incorrect Username or password" {
			return "Incorrect Username or password"
		}
		return user["token"]
	}
	return ""
}



func main()  {
	var Url string
	for i:=0;i<3 ;i++ {
		fmt.Println("This script get Edgescale login token:")
		fmt.Println("select 1、api.edgescale.org ")
		fmt.Println("select 2、test.api.edgescale.org ")
		fmt.Print("select request URL : ")
		input_number := bufio.NewScanner(os.Stdin) //初始化一个扫表对象
		input_number.Scan()
		line_number := input_number.Text() //把输入内容转换为字符串
		if line_number == "1" {
			fmt.Println("https://api.edgescale.org")
			Url = "https://api.edgescale.org"
		} else if line_number == "2" {
			fmt.Println("https://test.api.edgescale.org")
			Url = "https://test.api.edgescale.org"
		} else {
			fmt.Println("input number error")
			continue
		}

		fmt.Print("input user name: ")
		input := bufio.NewScanner(os.Stdin)//初始化一个扫表对象
		input.Scan()
		line := input.Text()//把输入内容转换为字符串

		fmt.Print("input user pwd: ")
		input2 := bufio.NewScanner(os.Stdin)//初始化一个扫表对象
		input2.Scan()
		line2 := input2.Text()//把输入内容转换为字符串

		v:=Token(Url,line,line2)
		fmt.Println("Result:")

		fmt.Println(v)

		if v !="Incorrect Username or password"{
			fmt.Println("Incorrect Username or password")
			break
		}
		fmt.Println(" ")
		//fmt.Println("======Please Enter again =========")
		//fmt.Println(" ")
	}

	fmt.Println(" ")
	fmt.Println("======Enter key is exit=========")
	input3:=bufio.NewScanner(os.Stdin)//初始化一个扫表对象
	input3.Scan()
}