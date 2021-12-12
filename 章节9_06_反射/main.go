package main

import (
	"reflect"
	"fmt"
)

type People struct {
	Name    string `xml:"name"`
	Address string
}

func main() {
	//a:=1.5
	//fmt.Println(reflect.TypeOf(a))
	//fmt.Println(reflect.ValueOf(a))

	/*
	获取结构体属性的值
	 */
	//peo:=People{"smallming","北京海淀"}
	//
	//v:=reflect.ValueOf(peo)
	//
	//fmt.Println(v.NumField())
	//fmt.Println(v.FieldByIndex([]int{0}))
	//
	//content:="address"
	//fmt.Println(v.FieldByName(content))

	/*
	设置属性的值
	 */
	//content := "Name"
	//peo := new(People)
	//v := reflect.ValueOf(peo).Elem()
	////需要修改属性的内容时,要求结构体中属性名首字母大写才可以设置
	//fmt.Println(v.FieldByName(content).CanSet())
	//v.FieldByName(content).SetString("smallming")
	//v.FieldByName("Address").SetString("北京海淀")
	//fmt.Println(peo)

	/*
	获取标记
	 */
	t := reflect.TypeOf(People{})
	fmt.Println(t.FieldByName("Name"))
	name, _ := t.FieldByName("Name")
	fmt.Println(name.Tag)
	fmt.Println(name.Tag.Get("xml"))
}
