package main
import (
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


func Token( myname string ,mypwd string)(tok string) {
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
		return tok
	}
	if resp.StatusCode() == 200 {
		body, err := resp.Body()
		if err != nil {
			log.Println(err)
			return tok
		}
		var user = map[string]string{}
		json.Unmarshal(body,&user)
		if user["message"] == "Incorrect Username or password" {
			return "Incorrect Username or password"
		}
		return user["token"]
	}
	return tok
}


func main()  {
		v:=Token("yuren.zhang@nxp.com","adminadmin")
		fmt.Println(v)
		p,err := os.Getwd()

		if err !=nil {
			fmt.Println(err)
		}
		pathfile:= p + "/EdgeScaleTestAuto/src/test/java/com/edgescale/common"
		fmt.Println(pathfile)
}