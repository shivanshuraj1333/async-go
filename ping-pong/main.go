package main

import (
	"fmt"
	"sync"
)
import "time"
import "os"

type Ball uint64

func Play(playerName string, table chan Ball, wg *sync.WaitGroup) {
	defer wg.Done()
	var lastValue Ball = 1
	for {
		ball := <- table // get the ball
		fmt.Println(playerName, ball)
		ball += lastValue
		if ball < lastValue { // overflow
			os.Exit(0)
		}
		lastValue = ball
		table <- ball // bat back the ball
		time.Sleep(time.Second)
	}
}

func main() {
	table := make(chan Ball)
	wg := new(sync.WaitGroup)
	go func() {
		table <- 1 // throw ball on table
	}()
	wg.Add(2)
	go Play("A:", table, wg)
	go Play("B:", table, wg)
	wg.Wait()
}