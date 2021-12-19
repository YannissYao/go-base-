package main

import (
	"container/list"
	"fmt"
)

type People struct {
	Name string
	Age  int
}

func main() {
	var peo People
	//fmt.Println(peo)
	//fmt.Printf("%p",&peo)

	peo = People{"smallming", 17}
	fmt.Println(peo)
	peo2 := People{Age: 17, Name: "smallming"}
	fmt.Println(peo2)

	//peo.Name = "名字"
	//peo.Age = 19
	//fmt.Println(peo.Name,peo.Age)

	fmt.Printf("%p %p\n", &peo, &peo2)
	fmt.Println(&peo == &peo2)
	list := list.New()
	list.PushBack(2)

	type aa interface {
		me(i int) People
	}
	type pp struct from aa {
	}
	}
