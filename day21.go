package main
import "fmt"

func main()  {
	var s = student{
		human:human{name:"教科书",age:2},
		book:"语文book",
	}
	fmt.Println(s.book)

}

