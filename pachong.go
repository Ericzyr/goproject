package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"github.com/PuerkitoBio/goquery"
	"reflect"
)


var a = 0
var url = "https://tieba.baidu.com/p/4608718475"
var isNext = true


func GetDetail(doc *goquery.Document){
	// 解析详情页函数

	doc.Find("img.BDE_Image[size]").Each(func(i int, s *goquery.Selection) {
		_picSrc, _exist := s.Attr("src")
		if (_exist) {
			fmt.Println(_picSrc)
			resp, err := http.Get(_picSrc)
			if err != nil {
				fmt.Println("访问图片出错")
			}

			_data ,_err2 := ioutil.ReadAll(resp.Body)
			if _err2 != nil {
				panic(_err2)
			}
			//保存到本地
			ioutil.WriteFile(fmt.Sprintf("./images/20190425_%d.jpg", a), _data, 0644)
			fmt.Println("图片下载成功")
			a ++
		}
	})
}

func GetMovie(startUrl string) (bool){
	fmt.Println(url)

	_doc, _err := goquery.NewDocument(startUrl)

	// 开始下载单页面的图片
	//GetDetail(_doc)
	if _err != nil {
		fmt.Println("下载请求错误")
	}
	//fmt.Println("sadf",reflect.ValueOf(*_doc))
	_next, _exist := _doc.Find("span.tP+a").Attr("href")
	// 组装下一页url
	url = "https://tieba.baidu.com" + _next
	fmt.Println(_next)
	fmt.Println(_exist)
	return _exist

}

func main() {
	for isNext {
		isNext = GetMovie(url)
		//fmt.Println(url)

	}
}