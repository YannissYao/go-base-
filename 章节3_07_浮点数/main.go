package main

import (
	"fmt"
	"math"
)

func main() {
	//整型和浮点型转换
	//var f float64 = 1
	//var i = 2
	//f = float64(i)
	//fmt.Println(f, i)
	maxInt64 := math.MaxInt32
	fmt.Println(maxInt64)

	//最大值
	//fmt.Println(math.Max
	//Float32)
	//fmt.Println(math.MaxFloat64)

	//浮点型之间相互转换
	//var a float32 = 1.5
	//var b float64 = 1.6
	//fmt.Println(a+float32(b))
	//fmt.Println(float64(a)+b)

	//数学运算结果
	var a, b = 3, 2
	var c, d float64 = 3, 2

	fmt.Println(a / b)
	fmt.Println(c / d)
}
