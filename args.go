package main
import("fmt";"os")

func main(){
	if len(os.Args)>1{
		fmt.Println(os.Args[1])
		fmt.Println(len(os.Args))
		param1:=os.Args[1]
		param2:=os.Args[2]
		fmt.Println(param1)
		fmt.Println(param2)
	} else {
		fmt.Println("需要参数")
		fmt.Println(len(os.Args))
	}
}