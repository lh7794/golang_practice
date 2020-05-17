package main

//声明日志写入器接口，写入器（日志写入设备）可以为命令行、文件、网络等
type LogWriter interface {
	Write(data interface{}) error
}

//日志器
type Logger struct {
	writeList []LogWriter
}

//注册一个日志写入器，将日志写入器的接口添加到writeList中

func (l *Logger) RegisterWriter(writer LogWriter) {
	l.writeList = append(l.writeList, writer)
}

//将一个data类型的数据写入日志

func (l *Logger) Log(data interface{}) {
	for _, writer := range l.writeList {
		writer.Write(data)
	}
}

//创建日志器的实例

func newLogger() *Logger {
	return &Logger{}
}
