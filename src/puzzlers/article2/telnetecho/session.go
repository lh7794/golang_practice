package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func handleSession(conn net.Conn, exitChan chan int) {
	fmt.Println("session started: ")

	//创建一个网络连接数据的读取器

	reader := bufio.NewReader(conn)

	//接受数据的循环

	for {

		//读取字符串，直到碰到回车返回
		str, err := reader.ReadString('\n')

		//数据读取正确

		if err == nil {

			//去掉字符串尾部的回车
			str = strings.TrimSpace(str)
			//处理Telnet指令

			if !processTelnetCommand(str, exitChan) {
				conn.Close()
				break
			}

			//Echo逻辑，发什么数据，原样返回
			conn.Write([]byte(str + "\r\n"))

		} else {
			//发生错误
			fmt.Println("Session closed")
			conn.Close()
			break

		}
	}
}
