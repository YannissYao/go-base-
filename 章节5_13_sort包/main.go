package main

import (
	"fmt"
	"sort"
)

func main() {
	//n := []int{1, 5, 3, 3, 4}
	//sort.Ints(n)
	//fmt.Println(n)
	//int类型切片排序
	num := []string{"a", "c", "b"}
	////sort.Ints(num)//升序排序
	sort.Strings(num)
	fmt.Println(num)
	//sort.Sort(sort.IntSlice(num))//升序排序
	//fmt.Println(num)
	//
	sort.Sort(sort.Reverse(sort.StringSlice(num))) //降序排序
	fmt.Println(num)

	//float类型切片排序
	//f := []float64{1.2,9.4,5.6,}
	//sort.Float64s(f)
	//fmt.Println(f)
	//
	//sort.Sort(sort.Reverse(sort.Float64Slice(f)))
	//fmt.Println(f)

	//string类型排序
	//s := []string{"a", "j", "a我们", "我们"}
	//sort.Strings(s)
	//fmt.Println(s)
	////切片应该是一个升序排序的切片
	//n := sort.SearchStrings(s, "a我们1")
	//if n < len(s) {
	//	if s[n] == "a我们1" {
	//		fmt.Println("存在")
	//	} else {
	//		fmt.Println("不存在")
	//	}
	//} else {
	//	fmt.Println("不存在")
	//}
	//
	//sort.Sort(sort.Reverse(sort.StringSlice(s)))
	//fmt.Println(s)

}
