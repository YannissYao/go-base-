package main

import (
	"fmt"
	"os/user"
)

func main() {
	//u,error:=user.Current()
	//u,error:=user.LookupId("S-1-5-21-2183371115-4052466699-1736786417-1001")
	u, error := user.Lookup("Joeysin")
	if error != nil {
		fmt.Println(error)
		return
	}
	fmt.Println(u.Uid)
	fmt.Println(u.Name)
	fmt.Println(u.Gid)
	fmt.Println(u.HomeDir)
	fmt.Println(u.Username)
	fmt.Println(u.GroupIds())
}
