package main
import (
	"encoding/json"
	"fmt"
)

type Ptersion struct {
	Name string
	Age int
}

func main() {
	mapt := make(map[int]string)
	mapt[1]="hos"
	mapt[2]="string"
	mapt[3] ="got"
	fmt.Println(mapt)
	maap := map[string]int{"sfst":12,"two":2}
	fmt.Println(maap)
	p:=Ptersion{"jcker",12}
	fmt.Println(p)
	b,err:=json.Marshal(p)
	if err !=nil{
		println("json error")
	}
	fmt.Println(string(b))
}