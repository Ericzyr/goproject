package main

import (
	"github.com/ctfang/command"
	"log"
)

func main() {
	app := command.New()
	app.AddCommand(Echo{})
	app.Run()
}
// Echo 需要实现接口 CommandInterface
type Echo struct {
}

func (Echo) Configure() command.Configure {
	return command.Configure{
		Name:"echo",
		Description:"示例命令 echo",
	}
}

func (Echo) Execute(input command.Input) {
	log.Println("echo command")
}