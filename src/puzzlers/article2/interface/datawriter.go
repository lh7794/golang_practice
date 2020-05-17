package main

import "fmt"

type DataWriter interface {
	WriteData(data interface{}) error
}

type file struct {
}

//实现DataWriter接口的WriteData方法
func (d *file) WriteData(data interface{}) error {

	//模拟写入数据
	fmt.Println("WriteData:", data)
	return nil
}

func newFile() *file {
	return &file{}
}

func main() {
	//实例化file
	//f := new(file) //f是指针类型
	f := newFile()
	f.WriteData("data")
	//f类型是DataWriter接口类型的实现形式
	var writer DataWriter = f //调用时取地址

	writer.WriteData("data")

	_, ok := interface{}(f).(DataWriter)
	fmt.Printf("*Cat implements interface Pet: %v\n", ok)
}
