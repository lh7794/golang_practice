package main

import "fmt"

func main() {
	//var m map[string]int 声明map会引发panic

	m := map[string]int{
		//"lihao": 7794, 初始化一个map 即使是nil 添加键元素对也不会引发panic
	}
	value := 444
	m["lihao1"] = value //值为nil的map中添加键-元素对会引发panic

	key := "two"
	elem, ok := m["two"]
	fmt.Printf("The element paired with key %q in nil map: %d (%v)\n",
		key, elem, ok)

	fmt.Printf("The length of nil map: %d\n",
		len(m))

	fmt.Printf("Delete the key-element pair by key %q...\n",
		key)
	//	delete(m, key)

	fmt.Println("m ", m)

	elem = 2
	fmt.Println("Add a key-element pair to a nil map...")
	//m["two"] = elem // 这里会引发panic。
}
