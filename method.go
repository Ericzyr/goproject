package  main
import ("fmt")

type fruit struct{
	name string
	age int
}

func (f fruit) drink()  {
	f.name="2banana"
	fmt.Println(f.name)
}


func main()  {
	a :=new(fruit)
	a.name="1apple"
	fmt.Println(a.name)
	a.drink()
}