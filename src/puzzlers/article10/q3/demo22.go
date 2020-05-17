package main

import "fmt"

func main() {
	ch1 := make(chan int, 3) //容量为0，为什么编译结果会发送0 1，在接收0 1.
	fmt.Println(cap(ch1))
	fmt.Println(len(ch1))
	// 发送方。
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("Sender: sending element %v...\n", i)
			ch1 <- i
		}
		fmt.Println("Sender: close the channel...")
		close(ch1)
	}()

	// 接收方。
	for {
		elem, ok := <-ch1
		if !ok {
			fmt.Println("Receiver: closed channel")
			break
		}
		fmt.Printf("Receiver: received an element: %v\n", elem)
	}

	fmt.Println("End.")
}
