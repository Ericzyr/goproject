package main

import "fmt"

func main()  {
	for i:=0;i<3 ;i++  {
		fmt.Println(&i)
		u :=&i
		fmt.Println(*u)
			defer func() {
				fmt.Println(&i)
	}()
	}
}

