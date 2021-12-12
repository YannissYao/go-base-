package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//var a, b float64 = 12.3, 13.5
	//fmt.Println(a, b)
	//
	//fmt.Println(math.Floor(a))
	//fmt.Println(math.Ceil(a))
	//fmt.Println(math.Abs(a))
	//fmt.Println(math.Modf(a))
	//fmt.Println(math.Max(a,b))
	//fmt.Println(math.Min(a,b))
	//fmt.Println(math.Pow(3,2))
	//fmt.Println(math.Round(b))

	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(10))
	fmt.Println(rand.Intn(10))
	fmt.Println(rand.Intn(10))
	fmt.Println(rand.Intn(10))
}
