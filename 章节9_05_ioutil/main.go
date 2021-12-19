package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//f, _ := os.Open("/Users/Joeysin/Desktop/032.txt")
	//b, _ := ioutil.ReadAll(f)
	//fmt.Println(string(b))

	b, _ := ioutil.ReadFile("/Users/Joeysin/Desktop/032.txt")
	fmt.Println(string(b))

	//ioutil.WriteFile("D:/go.txt", []byte("这是写的数据"), 0666)
	//fmt.Println("写入数据")

	//fi,_:=ioutil.ReadDir("D:/")
	//for _,n:=range fi{
	//	fmt.Println(n.Name())
	//}
}
