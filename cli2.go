package main
import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli"
	"os"
	"HttpRequest"
	"log"
)

const (
	URL = "https://api.edgescale.org";
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
		return user["message"]
	}
	return ""
}


func main()  {
	app := cli.NewApp()
	app.Name    = "Edgescale"                // 应用名称
	app.Usage   = "test cli"                 // 应用功能说明
	app.Author  = "Eric"                     // 应用作者
	app.Email   = "yuen.zhang@nxp.com"       // 邮箱
	app.Version = "2.0.1"                    // 版本


	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name : "name",             // 命令名称
			Usage : "user name",    	// 命令说明
			Value : "none",             // 默认值
		},
		cli.StringFlag{
			Name: "password,pwd",
			Value: "none",
			Usage: "number one for operation",
		},
	}
	app.Action  = func(c *cli.Context){
		name := c.String("name")
		pwd := c.String("password")
		fmt.Println(Token(name,pwd))
	}
	app.Run(os.Args)
}