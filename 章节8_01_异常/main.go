package main

import (
	"fmt"
)

func demo(i, k int) (r int, e error) {
	if k == 0 {
		//e= errors.New("除数不能为0")
		e = fmt.Errorf("%s %d %d", "除数不能为0,两个参数分别是:", i, k)
		return
	}
	r = i / k
	return
}

func main() {
	result, error := demo(6, 0)
	//fmt.Println(result,error)
	if error != nil{
		fmt.Println("程序执行出错,错误信息",error)
		return
	}
	fmt.Println("执行成功,结果为:", result)
}
