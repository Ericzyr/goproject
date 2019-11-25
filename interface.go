package main

import ("errors";"fmt")
//写一个接口
type USB interface {
	name() string
	Connecter
}
//这个是嵌入接口
type Connecter interface {
	connect()
}
//这是为接口写了一个结构类型
type applephoto struct {
	user string
}
//这个是接口的实现的方法 mothod
func (pc applephoto) name()string {
	return pc.user
}
//这个是接口的实现的方法 mothod
func (pc applephoto) connect() {
	fmt.Println("PC connect success",pc.user)
}

func main()  {
	var u USB = applephoto{"apple"}
	fmt.Println(u.name(),errors.New("elevator"))
	u.connect()
}
