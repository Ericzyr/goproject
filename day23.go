package main
import (
	"AddTest"
	"fmt"
)

func main() {
	num := AddTest.Add(2)
	fmt.Println("2 + 3 =",num(5))
}