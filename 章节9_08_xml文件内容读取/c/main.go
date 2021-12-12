package main

import (
	"encoding/xml"
	"io/ioutil"
	"fmt"
)

type Peoples struct {
	XMLName xml.Name `xml:"peoples"`
	Version string   `xml:"version,attr"`
	Peos    []People `xml:"people"`
}

type People struct {
	XMLName xml.Name `xml:"people"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Address string   `xml:"address"`
}

func main() {
	peo := new(Peoples)
	b, _ := ioutil.ReadFile("people.xml")
	xml.Unmarshal(b, peo)
	fmt.Println(peo)
}
