package main

import "fmt"

var BufferChan chan int
var MsgChan chan bool

func init() {
	BufferChan = make(chan int, 1)
	MsgChan = make(chan bool, 0)
}

func Produce() {
	for i := 0; i < 1000; i++ {
		fmt.Println("生产了:", i)
		BufferChan <- i
	}
}

func Consumption() {
	for i := 0; i < 1000; i++ {
		x := <-BufferChan
		fmt.Println("消费了:", x)
	}
	MsgChan <- true
}

func main() {
	go Produce()
	go Consumption()
	_ = <-MsgChan

}
