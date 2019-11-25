package main
import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// http.Get
func httpGet() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func httpPost() {
	resp, err := http.Post("http://www.baidu.com",
		"application/x-www-form-urlencode",
		strings.NewReader("name=abc")) // Content-Type post请求必须设置
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func main()  {
	httpGet()
	httpPost()
}