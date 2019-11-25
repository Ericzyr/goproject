package main
import (
	"encoding/json"
	"fmt"
	"HttpRequest"
	"log"
	"os"
)

const (
	URL = "https://api.edgescale.demo";
)

var (
	useragent ="Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.110 Safari/537.36";
	ContentType ="application/json"
)

func Token( myname string ,mypwd string)string {
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
		return user["token"]
	}
	return ""
}


func main()  {
	if len(os.Args)>1{
		v:=Token(os.Args[1],os.Args[2])
		fmt.Println(v)
	} else {
		fmt.Print("no parameter")
	}
}