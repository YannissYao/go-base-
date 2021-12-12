package people

type People struct {
	name string
	age  int
}

//设置name属性的方法
func (p *People) SetName(name string) {
	p.name = name
}

//获取name属性的方法
func (p *People) GetName() string {
	return p.name
}

func (p *People) GetAge() int {
	return p.age
}

func (p *People) SetAge(age int) {
	if age < 1 || age > 200 {
		age = 0
	} else {
		p.age = age
	}
}
