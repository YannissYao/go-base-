package main

import (
	"container/list"
	"fmt"
)

func main() {
	//实例化list
	mylist := list.New()
	fmt.Println(mylist)
	//添加
	mylist.PushFront("a")                   //["a"]
	mylist.PushBack("b")                    //["a","b"]
	mylist.PushBack("c")                    //["a","b","c"]
	mylist.InsertBefore("d", mylist.Back()) //["a","b","d","c"]
	mylist.InsertAfter("e", mylist.Front()) //[a,e,b,d,c]

	//遍历
	for e := mylist.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println("")

	//fmt.Println(mylist.Front().Value)
	//fmt.Println(mylist.Back().Value)

	//n := 5
	//var curr *list.Element
	//if n > 0 && n <= mylist.Len() {
	//	if n == 1 {
	//		curr = mylist.Front()
	//	} else if n == mylist.Len() {
	//		curr = mylist.Back()
	//	} else {
	//		curr = mylist.Front()
	//		for i := 1; i < n; i++ {
	//			curr = curr.Next()
	//		}
	//	}
	//} else {
	//	fmt.Println("n的数值不对")
	//}

	//fmt.Println("取出第n个:", curr.Value)

	//移动元素的顺序

	//mylist.MoveBefore(mylist.Front(),mylist.Back())
	//mylist.MoveAfter(mylist.Front(),mylist.Back())
	//mylist.MoveToFront(mylist.Back())
	//mylist.MoveToBack(mylist.Front())

	mylist.Remove(mylist.Front())
	//遍历
	for e := mylist.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println("")
}
