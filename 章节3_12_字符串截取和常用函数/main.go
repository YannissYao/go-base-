package main

import (
	"fmt"
	"strings"
)

func main() {
	s:="smallming张"
	fmt.Println(len(s))

	//截取子字符串

	//s1:=s[:5]
	//fmt.Println(s1)
	//fmt.Printf("%T\n",s1)

	//s2:=fmt.Sprintf("%c",s1)
	//fmt.Println(s2)
	//fmt.Printf("%T",s2)

	//字符长度和获取字符
	//arr:=[]rune(s)
	//fmt.Println(len(arr))
	//fmt.Println(arr[9])

	//for
	//for _,n:= range s{
	//	fmt.Printf("%c\n",n)
	//}

	fmt.Println(strings.Index(s,"l"))
	fmt.Println(strings.LastIndex(s,"l"))
	fmt.Println(strings.HasPrefix(s,"small张"))
	fmt.Println(strings.HasSuffix(s,"sming张"))
	fmt.Println(strings.Contains(s,"mingsmall"))
	fmt.Println(strings.ToLower(s))
	fmt.Println(strings.ToUpper(s))
	fmt.Println(strings.Replace(s,"m","k",-1))
	fmt.Println(strings.Repeat(s,10))
	fmt.Println(strings.Trim(s," "))
	fmt.Println(strings.TrimSpace(s))
	a:=strings.Split(s,"m")
	fmt.Println(a)
	fmt.Printf("%T\n",a)

	b:=[]string{"a","b","c"}
	fmt.Println(strings.Join(b,""))
}
