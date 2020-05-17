package main

import "fmt"

//创建日志器

func createLogger() *Logger {

	//创建日志器
	l := newLogger()

	//创建命令行写入器
	//cw := &consoleWriter{}
	cw := newConsolWriter()

	//注册命令行写入器到日志器中
	l.RegisterWriter(cw)

	//创建文件写入器
	fw := newFileWriter()

	//设置文件名

	if err := fw.SetFile("log.log"); err != nil {
		fmt.Println(err)
	}

	//注册文件写入器到日志器中

	l.RegisterWriter(fw)

	return l

}

func main() {
	l := createLogger()

	l.Log("hello")
}
