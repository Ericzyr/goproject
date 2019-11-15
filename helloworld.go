package main
import "fmt"
var age int = 12 + 3
var fruit string= "apples" + "oranges"

func main() {
	fmt.Println("Hello, world")
	fmt.Println("fist a go");
	fmt.Println(age)
	fmt.Println(fruit)
	tdsaf()
	}


func tdsaf() {
	fmt.Println("Hello, World!")
	fmt.Println("fist a go");fmt.Println(age)
	fmt.Println(fruit)
	var tot int	=max(2,5)
	fmt.Println(tot)
}

func max(num1, num2 int) int {
	/* 定义局部变量 */
	var result int

	if (num1 > num2) {
		result = num1
	} else {
		result = num2
	}
	return result
}


func forilf(){
	for true{
		fmt.Print("ok")
	}
}
