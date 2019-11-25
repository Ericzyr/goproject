package main

import (
	"fmt"
	"io"
	"os"
)

//写入文件create


func createfile(){
	fp,err:=os.Create("a.txt")
	if err != nil{
		fmt.Print("file create error")
		return
	}
	fp.WriteString("hello world \r\n")
	fp.WriteString("hello world \r\n")
	defer fp.Close()
}

func openfile()  {
	//这种打开是只读的方式打开的
	fp,err:=os.Open("a.txt")
	if err !=nil {
		fmt.Println(err)
	}
	arr:= make([]byte,1024*2)
	r,_err:=fp.Read(arr)
	if _err !=nil{
		fmt.Println(_err)
	}
	fmt.Println(string(arr[0:r]))
	defer fp.Close()
}

func openAndWriteFile()  {
	fp,err:=os.OpenFile("a.txt",os.O_RDWR,6)
	if err !=nil{
		fmt.Println(err)
	}
	count,_:= fp.Seek(0,io.SeekEnd)

	fp.WriteAt([]byte("my read Englist"),count)

	defer fp.Close()

}

func main()  {
	openAndWriteFile()
}