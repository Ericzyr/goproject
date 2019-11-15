
package main

import (
	"encoding/json"
	"fmt"
	"github.com/kirinlabs/HttpRequest"
	"log"
)

const URL = "https://test.api.edgescale.org"


var token ="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Inl1cmVuLnpoYW5nQG54cC5jb20iLCJhZG1pbiI6MSwidWlkIjozMDYsImV4cCI6MTU2NDU2MTUxMiwiYWNjb3VudF90eXBlX2lkIjo0fQ.HqQ3-L2xfDaVT8ASMezf-vTcrZCUO9QMwjS5WI9mTH0"

var useragent ="Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.110 Safari/537.36"

var ContentType ="application/json"

func main() {

	req := HttpRequest.NewRequest()

	// 设置超时时间，不设置时，默认30s
	req.SetTimeout(5)

	// 设置Headers
	req.SetHeaders(map[string]string{
		"user-agent": useragent,
		"Content-Type": ContentType,
		"dcca_token": token,
	})

	var postData = map[string]interface{}{
		"offset": "0",
		"limit": "10",
	}

	// GET 默认调用方法
	resp, err := req.Get(URL + "/v1/customers",postData)

	// GET 传参调用方法
	// 第2个参数默认为nil，也可以传参map[string]interface{}
	// 第2个参数不为nil时，会把传入的map以query传参的形式重新构造新url
	// 新的URL: http://127.0.0.1:8000?name=flyfreely&id=1&title=csdn

	//resp, err := req.Get("http://127.0.0.1:8000?name=flyfreely", postData)

	// POST 调用方法

	//resp, err := req.Post("http://127.0.0.1:8000", postData)

	if err != nil {
		log.Println(err)
		return
	}
	if resp.StatusCode() == 200 {
		body, err := resp.Body()
		if err != nil {
			log.Println(err)
			return
		}

		var user map[string][]map[string] interface{}
		json.Unmarshal(body,&user)
		fmt.Println(string(body))
		fmt.Println(user["customers"][0]["id"])

	}
}