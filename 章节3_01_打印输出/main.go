package main

import (
	"fmt"
	//"os"
)

func main() {
	//fmt.Fprintln(os.Stdout,"内容1","内容2")
	//fmt.Fprintln(os.Stdout,"内容3","内容4")
	//fmt.Fprint(os.Stdout,"内容5","内容6")
	//fmt.Fprint(os.Stdout,"内容7","内容8","内容9")

	//fmt.Println("1","2")
	//fmt.Print("3","4")
	//fmt.Println("5")
	//fmt.Println("6")


	a:=fmt.Sprintln("1","2")
	fmt.Println(a)
	b:=fmt.Sprint("3","4","5")
	fmt.Println(b)
	fmt.Sprintln("4")
}
