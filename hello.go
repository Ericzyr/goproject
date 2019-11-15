package main
import "fmt"
type persions struct {
	name string
	age int
}
func AB(per *persions){
	per.name="tom"
	per.age= 19
	fmt.Println("AB",per)
}
func main(){
	a:= persions{
		name: "jack",
		age : 10,
	}
	fmt.Println("a",a)
	AB(&a)
	fmt.Println("a",a)
}
