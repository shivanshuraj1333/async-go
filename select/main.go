package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func(_ch chan int){
		time.Sleep(time.Millisecond*4)
		_ch <- 5
	}(ch)
	select {
		case x :=<- ch :
			fmt.Println(x)
		case <- time.After(time.Millisecond*5):
			fmt.Print("Timeout!")
	}
}
