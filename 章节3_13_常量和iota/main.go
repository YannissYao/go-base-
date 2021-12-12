package main

import "fmt"

func main() {

	//单个常量定义

	//const a string = "smallming"
	//const b = 1
	//const c = 1.5
	//
	//fmt.Printf("%T %T \n",b,c)
	//
	//fmt.Println(b+c)

	//i1 := 1
	//i2 := 1.5
	//fmt.Println(i1+i2)

	//const d = 1*2+3
	//i:=2
	//const e = 1*i +3

	/*
	定义常量组
	 */
	//const (
	//	a=5
	//	b=6
	//	c=7
	//)
	//fmt.Println(a,b,c)

	/*
	常量生成器
	 */
	const (
		a = iota
		b
		c
	)
	fmt.Println(a, b, c)

	const (
		d = iota + 8
		e
		f
	)
	fmt.Println(d, e, f)

	const (
		g = 10   //iota = 0
		h = iota // = 1
		i        // 2
	)
	fmt.Println(g, h, i)

	const (
		j = iota //iota = 0
		k        //iota=1
		l = 6    //iota =2
		m        //iota=3
		n = iota //iota=4
		o        //iota=5
	)
	fmt.Println(j, k, l, m, n, o)
}
