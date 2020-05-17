package main

import (
	"fmt"
	"strings"
)

func processTelnetCommand(str string, exitChan chan int) bool {

	//@close指令表示终止本次会话

	if strings.HasPrefix(str, "@close") {
		fmt.Println("Session closed")
		return false
		//@shutdown指令表示终止服务进程
	} else if strings.HasPrefix(str, "@shutdown") {
		fmt.Println("Server shutdown")

		//往通道中写入0，阻塞等待接收方处理
		exitChan <- 0

		return false
	}
	//打印输入的字符串
	fmt.Println(str)

	return true
}
