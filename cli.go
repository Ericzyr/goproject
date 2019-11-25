package main
import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)
func main()  {
	app := cli.NewApp()
	app.Name    = "Edgescale"                       // 应用名称
	app.Usage   = "test cli"                   // 应用功能说明
	app.Author  = "Eric"                      // 应用作者
	app.Email   = "yuen.zhang@nxp.com"       // 邮箱
	app.Version = "2.0.1"                       // 版本


	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name : "name",             // 命令名称
			Usage : "user name",    	 // 命令说明
			Value : "none",             // 默认值
		},
		cli.StringFlag{
			Name: "age",
			Value: "none",
			Usage: "number",
		},
	}
	app.Action  = func(c *cli.Context) error {
		name := c.String("name")
		age := c.String("age")
		fmt.Println("name is :",name)
		fmt.Println("age is :",age)
		return nil
	}
	app.Run(os.Args)
}