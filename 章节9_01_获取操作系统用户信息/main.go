package main

import (
	"os/user"
	"fmt"
)

func main() {
	//u,error:=user.Current()
	//u,error:=user.LookupId("S-1-5-21-2183371115-4052466699-1736786417-1001")
	u,error:=user.Lookup("LAPTOP-M7D47U95\\zhang")
	if error!=nil{
		fmt.Println(error)
		return
	}
	fmt.Println(u.Uid)
	fmt.Println(u.Name)
	fmt.Println(u.Gid)
	fmt.Println(u.HomeDir)
	fmt.Println(u.Username)
}