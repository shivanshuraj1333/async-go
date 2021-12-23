package main

import (
	"fmt"
	"time"
)

var (
	requestIDs = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12 ,13, 14, 15}
)

var concurrencyLevel = 3

func makeRequest(ID int) {
	time.Sleep(time.Second)
	fmt.Println(ID)
}

func main() {
	queue := make(chan bool, concurrencyLevel)

	for _, _ID := range requestIDs {
		queue <- true
		go func (ID int) {
			defer func() {
				<- queue
			}()
			makeRequest(ID)
		}(_ID)
	}
}
