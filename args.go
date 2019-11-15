package main
import("fmt";"os")

func main(){
	if len(os.Args)>1{
		param1:=os.Args[1]
		fmt.Println(param1)
	} else {
		fmt.Println("需要参数")
		fmt.Println(len(os.Args))
	}
}