package main
import (
	"bufio"
	"encoding/json"
	"fmt"
	"HttpRequest"
	"log"
	"os"
)

const (
	URL = "https://api.edgescale.org"
)

var (
	useragent ="Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.110 Safari/537.36";
	ContentType ="application/json"
)

func Token(myname string ,mypwd string)string {
	req := HttpRequest.NewRequest()
	req.SetHeaders(map[string]string{
		"user-agent": useragent,
		"Content-Type": ContentType,
	})

	var postData = map[string]interface{}{
		"username": myname,
		"password": mypwd,
	}
	resp, err := req.Post(URL + "/v1/users/login",postData)

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
	for i:=0;i<3 ;i++  {

		fmt.Println("This script get Edgescale login token:")
		fmt.Print("input user name: ")
		input := bufio.NewScanner(os.Stdin)//初始化一个扫表对象
		input.Scan()
		line := input.Text()//把输入内容转换为字符串

		fmt.Print("input user pwd: ")
		input2 := bufio.NewScanner(os.Stdin)//初始化一个扫表对象
		input2.Scan()
		line2 := input2.Text()//把输入内容转换为字符串

		v:=Token(line,line2)
		fmt.Println("Result:")

		fmt.Println(v)

		if v !="Incorrect Username or password"{
			break
		}
		fmt.Println(" ")
		fmt.Println("======Please Enter again =========")
		fmt.Println(" ")
	}

	fmt.Println(" ")
	fmt.Println("======Enter key is exit=========")
	input3:=bufio.NewScanner(os.Stdin)//初始化一个扫表对象
	input3.Scan()
}