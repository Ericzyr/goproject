package main

import (
	"crypto/md5"
	"crypto/x509"
	"encoding/base64"
	"fmt"
)

func main() {
	Md5Inst:= md5.New()
	Md5Inst.Write([]byte("在不同的情况下"))
	Result := Md5Inst.Sum([]byte("1r"))
	fmt.Printf("%x\n\n",Result)
	fmt.Print(base64.StdEncoding.EncodeToString(Result))

	priv, err := x509.ParsePKCS1PrivateKey(Result)
	fmt.Println(priv)
	fmt.Println(err)
}
