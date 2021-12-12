package main

import "fmt"

func demo(name string,hover ... string){
	fmt.Println(name,"的爱好是")
	for i,n := range hover{
		fmt.Println(i,n)
	}
}

func main() {
	demo("张三","看书","写代码","看佳明哥视频")
}
