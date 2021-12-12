package main

import "fmt"

func main() {
	num := [5]int{1,2,3,4,5}
	s:=num[0:2]
	fmt.Println(s)
	fmt.Printf("%T\n",s)
	fmt.Printf("%T\n",num)

	fmt.Printf("%p %p\n",s,&num[0])
	s[0]=6
	fmt.Println(num,s)

	s= append(s,7)
	fmt.Println(num,s)

	s= append(s,8,9,10)
	fmt.Println(num,s)
	fmt.Printf("%p %p\n",s,&num[0])

	fmt.Println(s)
	//删除代码
	n:=2//删除元素的索引
	newSlice := s[0:n]
	newSlice = append(newSlice,s[n+1:]...)
	fmt.Println(s)//导致原切片内容也跟随变化,不要使用原切片
	fmt.Println(newSlice)

	fmt.Printf("%p %p",s,newSlice)
}
