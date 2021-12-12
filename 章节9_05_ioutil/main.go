package main

import (
	"io/ioutil"
	"fmt"
)

func main() {
	//f,_:=os.Open("D:/go.txt")
	//b,_:=ioutil.ReadAll(f)
	//fmt.Println(string(b))

	//b, _ := ioutil.ReadFile("D:/go.txt")
	//fmt.Println(string(b))

	//ioutil.WriteFile("D:/go.txt", []byte("这是写的数据"), 0666)
	//fmt.Println("写入数据")

	fi,_:=ioutil.ReadDir("D:/")
	for _,n:=range fi{
		fmt.Println(n.Name())
	}
}
