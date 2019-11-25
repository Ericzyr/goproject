package main

import (
	"fmt"
	"goquery"
	"log"
	"os"
)

func ExampleScrape() {
	file :="news.txt"
	fout ,err:=os.Create(file)
	defer fout.Close()
	if err!=nil{
		fmt.Println(file,err)
		return
	}
	doc, err := goquery.NewDocument("https://www.baidu.com")
	if err != nil {
		log.Fatal(err,"\n")
	}
	doc.Find("tbody").Each(func(i int, s *goquery.Selection) {
		fmt.Println(s.Text())
		fout.WriteString(s.Text())
	})
}
func main() {
	ExampleScrape()
}