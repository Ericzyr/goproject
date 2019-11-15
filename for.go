package main
import "fmt"
func Number(args int)  int  {
	for i :=1;i <=args ; i++ {
		fmt.Println(i)
	}
	return 4
}

func paper(){
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

func main(){
	Number(10)
	paper()
}