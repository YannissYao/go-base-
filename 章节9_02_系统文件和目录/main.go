package main

import (
	"fmt"
	"os"
)

func main() {
	//err := os.Mkdir("/Users/Joeysin/Desktop/aa", os.ModeDir)
	//if err != nil {
	//	fmt.Println("文件夹创建失败,", err)
	//	return
	//}
	//fmt.Println("文件夹创建成功")

	//err := os.MkdirAll("D:/godir/a/b", os.ModeDir)
	//if err != nil {
	//	fmt.Println("文件夹创建失败,", err)
	//	return
	//}
	//fmt.Println("文件夹创建成功")

	//f,err:=os.Create("D:/godir/test.txt")
	//if err!=nil{
	//	fmt.Println("文件创建失败,",err)
	//	return
	//}
	//fmt.Println("文件创建成功,",f)

	//err:=os.Rename("D:/godir","D:/godir1")
	//if err!=nil{
	//	fmt.Println("重命名失败",err)
	//	return
	//}
	//fmt.Println("文件夹重命名成功")

	//err:=os.Rename("D:/godir1/test.txt","D:/godir1/test1.txt")
	//if err!=nil{
	//	fmt.Println("重命名失败",err)
	//	return
	//}
	//fmt.Println("重命名成功")

	//f, err := os.Open("/Users/Joeysin/Desktop/032.txt")
	//if err != nil {
	//	fmt.Println("文件获取失败", err)
	//	return
	//}
	//fileInfo, err := f.Stat()
	//if err != nil {
	//	fmt.Println("文件信息获取失败", err)
	//	return
	//}
	//fmt.Println(fileInfo.Size())
	//fmt.Println(fileInfo.ModTime())
	//fmt.Println(fileInfo.Mode())
	//fmt.Println(fileInfo.IsDir())
	//fmt.Println(fileInfo.Name())

	err := os.RemoveAll("/Users/Joeysin/Desktop/aa")
	if err != nil {
		fmt.Println("删除失败,", err)
		return
	}
	fmt.Println("删除成功")
}
