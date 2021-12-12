package main

import "fmt"

func main() {
	//使用make实例化
	//var m map[string]string=make(map[string]string)
	//fmt.Println(m==nil)
	//fmt.Printf("%p",m)

	//实例化时赋初始值
	//m := map[string]string{"张三": "1388888888", "李四": "13999999999"}
	//fmt.Println(m)
	//
	////key不存在表示新增
	//m["王五"] = "13555555555"
	//fmt.Println(m)
	////key存在覆盖之前元素
	//m["李四"] = "13444444444"
	//fmt.Println(m)
	//
	////删除是,key存在表示删除
	//delete(m, "李四")
	//fmt.Println(m)
	//delete(m, "赵六")
	//fmt.Println(m)
	//
	////查看key对应的value
	////如果key存在,返回key对应的value
	//fmt.Println(m["张三"])
	////如果key不存在,返回value类型的默认值
	//fmt.Println(m["赵六"])

	//value:key对应的值,ok表示key是否存在
	//value, ok := m["李四"]
	//if ok {
	//	fmt.Println(value, ok)
	//}
	m2 := map[string]string{}
	fmt.Println(m2 == nil)
	//循环遍历map所有内容
	//for key, value := range m {
	//	fmt.Println(key, value)
	//}
}
