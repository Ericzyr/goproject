package main
import "fmt"

type human struct {
	name string
	age int
}

type student struct {
	human
	book string
}

func main()  {
	var s = student{
		human:human{name:"教科书",age:2},
		book:"语文book",
	}
	fmt.Println(s.book)

}