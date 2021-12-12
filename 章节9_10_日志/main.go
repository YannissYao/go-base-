package main

import (
	"log"
	"os"
)

func main() {
	//log.Println("打印日志信息")
	//log.Panicln("打印panic日志信息")
	//log.Fatal("打印fatal信息")
	f, _ := os.OpenFile("D:/golog.log", os.O_APPEND|os.O_CREATE, 0777)
	logger := log.New(f, "[Info]", log.Ltime)
	logger.Println("打印日志信息123\r\n")
}
