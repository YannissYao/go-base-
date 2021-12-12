package main

import (
	"time"
	"fmt"
)

func main() {
	//声明时间变量
	//var t time.Time
	//fmt.Println(t)

	//获取当前时间
	//t := time.Now()
	//fmt.Println(t)

	//通过纳秒时间戳创建时间变量
	//t1:=time.Unix(0, t.UnixNano())
	//fmt.Println(t1)

	//自己指定时间
	//t:=time.Date(2022,2,4,18,8,9,123456,time.Local)
	//fmt.Println(t)
	//fmt.Println(t.Year())
	//fmt.Println(int(t.Month()))
	//fmt.Println(t.Day())
	//y,m,d := t.Date()
	//fmt.Println(y,m,d)
	//fmt.Println(t.Hour())
	//fmt.Println(t.Minute())
	//fmt.Println(t.Second())
	//fmt.Println(t.Clock())
	//fmt.Println(t.Nanosecond())
	//
	//fmt.Println(t.Unix())//秒差,举例1970年1月1日秒差
	//fmt.Println(t.UnixNano())


	//时间向string转换
	//t:= time.Now()
	//s:=t.Format("2006/01/02 15:04:05")
	//fmt.Println(s)

	//string向时间转换
	s:="2022-02-04 18:08:09"
	t,_:=time.Parse("2006-01-02 15:04:05",s)
	fmt.Println(t)

}
