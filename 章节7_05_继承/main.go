package main

import "fmt"

//匿名属性
//type People struct {
//	string
//	int
//}
//
//
//func main() {
//	peo:=People{"张三",16}
//	fmt.Println(peo)
//	peo.string="李四"
//	peo.int=17
//	fmt.Println(peo.string,peo.int)
//}

type People struct {
	name string
	age int
}

type Teacher struct {
	classroom string
	People
}

func main() {
	tea:=Teacher{"302班级",People{"张三",17}}
	fmt.Println(tea.classroom,tea.name,tea.age)
}
