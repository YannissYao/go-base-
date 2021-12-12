package main

import (
	"os"
	"fmt"
)

func main() {
	filePath := "D:/go.txt"
	f,err:=os.OpenFile(filePath,os.O_APPEND,0660)
	defer f.Close()
	if err!=nil{
		f,_=os.Create("D:/go.txt")
	}
	f.Write([]byte("输入内容123"))
	f.WriteString("\r\n再次\t输入内容")
	fmt.Println("程序结束")
}
