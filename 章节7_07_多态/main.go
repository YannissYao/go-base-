package main

import "fmt"

//type Live interface {
//	run()
//	eat()
//}
//
//type People struct {
//	name string
//}
//
//func (p *People) run(){
//	fmt.Println(p.name,"在跑步")
//}
//func (p People) eat(){
//	fmt.Println(p.name,"在吃饭")
//}
//func main() {
//	var live Live = &People{"张三"}
//	live.eat()
//	live.run()
//}

type Live interface {
	run() (int, error)
}
type People struct{}
type Animate struct{}

func (p *People) run() (int, error) {
	fmt.Println("人在跑步")
	return 1, nil
}
func (a *Animate) run() (int, error) {
	fmt.Println("动物在奔跑")
	return 1, nil
}

func allrun(live Live) {
	live.run()

}

func main() {

	peo := &Animate{}
	allrun(peo)
}
