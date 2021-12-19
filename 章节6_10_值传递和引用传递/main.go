package main

import "fmt"

func main() {
	a := 119
	b := "smallming"
	c := []int{1, 2}
	d := 119
	demo(&a, b, c, &d)
	fmt.Println(a, b, c, d)
}

func demo(i *int, s string, arr []int, content *int) {
	*i = 110
	s = "佳明哥"
	arr[0] = 3
	arr[1] = 4
	*content = 110
}
