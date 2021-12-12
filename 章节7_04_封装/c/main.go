package main

import (
	"people"
	"fmt"
)

func main() {
	peo:=new(people.People)
	peo.SetName("张三")
	fmt.Println(peo.GetName())
	peo.SetAge(500)
	if peo.GetAge()==0{
		fmt.Println("年龄设置不正确")
	}else{
		fmt.Println(peo.GetAge())
	}
}
