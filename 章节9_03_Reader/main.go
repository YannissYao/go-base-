package main

import (
	"os"
	"fmt"
)

func main() {
	//r:=strings.NewReader("hello 世界")
	//
	//b:=make([]byte,r.Size())
	//
	//n,err:=r.Read(b)
	//if err!=nil{
	//	fmt.Println("读取流数据失败,",err)
	//	return
	//}
	//fmt.Println("读取数据长度为",n)
	//fmt.Println("数据内容为:",string(b))

	f, _ := os.Open("D:/go.txt")
	fileInfo, _ := f.Stat()

	b := make([]byte, fileInfo.Size())

	f.Read(b)

	fmt.Println("文件中内容:", string(b))
}
