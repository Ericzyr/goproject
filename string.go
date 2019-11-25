package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main(){
	var strTest string = " http://www.cnblog.com/kaichenkai "

	// 是否以 http:// 开头
	ret1 := strings.HasPrefix(strTest, "http://")
	fmt.Println("HasPrefix result:", ret1)  // HasPrefix result: false

	// 是否以 .com 结尾
	ret2 := strings.HasSuffix(strTest, ".com")
	fmt.Println("HasSuffix result:", ret2)  // HasSuffix result: false

	// 返回 k 在字符串中首次出现的位置，没有则返回 -1
	ret3 := strings.Index(strTest, "k")
	fmt.Println("Index is:", ret3)  // Index is: 23

	// 返回 k 在字符串中最后出现的位置，没有则返回 -1
	ret4 := strings.LastIndex(strTest, "k")
	fmt.Println("LastIndex is:", ret4)  // LastIndex is: 30

	// 将 kai 换成 空格，替换1次，返回操作后的结果字符串
	ret5 := strings.Replace(strTest, "kai", " ", 1)
	fmt.Println("replace complete, strTest is:", ret5)  // replace complete, strTest is:  http://www.cnblog.com/ chenkai

	// 统计子字符串 kai 的出现次数
	ret6 := strings.Count(strTest, "kai")
	fmt.Println("Count is:", ret6)  // Count is: 2

	// 将字符串重复 n 次，返回操作后的结果字符串
	ret7 := strings.Repeat(strTest, 2)
	fmt.Println("Repeat result is:", ret7)  // Repeat result is:  http://www.cnblog.com/kaichenkai  http://www.cnblog.com/kaichenkai

	// 转小写
	ret8 := strings.ToLower(strTest)
	fmt.Println("Lower result is:", ret8)  // Lower result is:  http://www.cnblog.com/kaichenkai

	// 转大写
	ret9 := strings.ToUpper(strTest)
	fmt.Println("Upper result is:", ret9)  // Upper result is:  HTTP://WWW.CNBLOG.COM/KAICHENKAI

	// 去掉收尾空白字符，返回操作后的结果字符串
	ret10 := strings.TrimSpace(strTest)
	fmt.Println("TrimSpace:", ret10)  // TrimSpace: http://www.cnblog.com/kaichenkai

	// 去掉首尾指定字符，返回操作后的结果字符串
	ret11 := strings.Trim(strTest, "http://")
	fmt.Println("Trim:", ret11)  // Trim:  http://www.cnblog.com/kaichenkai

	// 去掉左侧指定字符，返回操作后的结果字符串
	ret12 := strings.TrimLeft(strTest, " http://")
	fmt.Println("TrimLeft:", ret12)  // TrimLeft: www.cnblog.com/kaichenkai

	// 去掉右侧指定字符，返回操作后的结果字符串
	ret13 := strings.TrimRight(strTest, "kai ")
	fmt.Println("TrimRight:", ret13)  // TrimRight:  http://www.cnblog.com/kaichen

	// 以空格分割，返回子串的 slice
	ret14 := strings.Fields(strTest)
	fmt.Println("Fields:", ret14)  // Fields: [http://www.cnblog.com/kaichenkai]

	// 以 .（点）进行分割，返回子串的 slice
	ret15 := strings.Split(strTest, ".")
	fmt.Println("Split:", ret15)  // Split: [ http://www cnblog com/kaichenkai ]

	// 以 $ 将 slice 中的元素进行拼接
	ret16 := strings.Join(ret15, "$")
	fmt.Println("Join:", ret16)  // Join:  http://www$cnblog$com/kaichenkai

	// 将整数转换为字符串
	ret17 := strconv.Itoa(1000)
	fmt.Println("Itoa:", ret17)  // Itoa: 1000

	// 将字符串转换为数字，前提是字符串是纯数字组成，不然会报错
	ret18, error := strconv.Atoi("100")
	if error == nil{
		fmt.Println("Atoi:", ret18)  // Atoi: 100
	}
}