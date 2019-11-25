package main

import (
	"fmt"
	"os"
	"github.com/mitchellh/cli"
)

func main() {
	c := cli.NewCLI("app", "1.0.0")    //这里指定了APP名字和版本
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"food": fooCommandFactory,    //定义foo命令和工厂
		"bar": barCommandFactory,    //定义bar命令和工厂
	}
	exitStatus, err := c.Run()
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(exitStatus)
}
//foo命令工厂
func fooCommandFactory() (cli.Command, error) {
	fmt.Println("I am foo command")
	return new(FooCommand), nil
}

//bar命令工厂
func barCommandFactory() (cli.Command, error) {
	fmt.Println("I am bar command")
	return new(BarCommand), nil
}


//foo命令实现
type FooCommand struct{

}

func (f *FooCommand) Help() string {
	return "help foo"
}
func (f *FooCommand) Run(args []string) int {
	fmt.Println("Foo Command is running")
	return 1
}
func (f *FooCommand) Synopsis() string {
	return "foo command Synopsis"
}


type BarCommand struct{}
func (b *BarCommand) Help() string {
	return "help bar"
}

func (b *BarCommand) Run(args []string) int {
	fmt.Println("bar Command is running")
	return 1
}
func (b *BarCommand) Synopsis() string {
	return "bar command Synopsis"
}