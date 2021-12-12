package main

type A struct {
	name string
}

func main() {
	a := A{name: "aa"}
	arr := [2]A{}
	arr[0] = a
	a.name = "bb"
	arr[1] = a
}
