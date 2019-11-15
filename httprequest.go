package main
import (
	"encoding/json"
	"fmt"
	"github.com/kirinlabs/HttpRequest"
	"log"
)
const URL = "https://api.edgescale.org"
const UserName  = "yuren.zhang@nxp.com"
const Passwd  = "adminadmin"

var useragent ="Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.110 Safari/537.36"
var ContentType ="application/json"

func Token()string {
	req := HttpRequest.NewRequest()
	req.SetHeaders(map[string]string{
		"user-agent": useragent,
		"Content-Type": ContentType,
	})

	var postData = map[string]interface{}{
		"username": UserName,
		"password": Passwd,
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



func main() {
	req := HttpRequest.NewRequest()

	req.SetHeaders(map[string]string{
		"user-agent": useragent,
		"Content-Type": ContentType,
		"dcca_token": Token(),
	})

	var postData = map[string]interface{}{
		"offset": "0",
		"limit": "10",
	}

	resp, err := req.Get(URL + "/v1/customers",postData)
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

		var user map[string][]map[string]interface{}
		json.Unmarshal(body, &user)
		fmt.Println(string(body))
		fmt.Println(user["customers"][0]["id"])
	}
}