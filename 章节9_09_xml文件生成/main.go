package main

import (
	"encoding/xml"
	"io/ioutil"
	"fmt"
)

type People struct {
	XMLName xml.Name `xml:"people"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Address string   `xml:"address"`
}

func main() {
	peo := People{Id: 123, Name: "smallming", Address: "北京海淀"}
	b, _ := xml.MarshalIndent(peo,"","	")
	b=append([]byte(xml.Header),b...)
	ioutil.WriteFile("D:/people.xml", b, 0777)
	fmt.Println("执行结束")
}
