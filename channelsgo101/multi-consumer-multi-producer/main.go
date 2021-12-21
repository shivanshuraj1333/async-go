package main

import (
	"fmt"
	"strconv"
	"sync"
)

var messages = [][]string{
	{
		"The world itself's",
		"just one big hoax.",
		"Spamming each other with our",
		"running commentary of bullshit,",
	},
	{
		"but with our things, our property, our money.",
		"I'm not saying anything new.",
		"We all know why we do this,",
		"not because Hunger Games",
		"books make us happy,",
	},
	{
		"masquerading as insight, our social media",
		"faking as intimacy.",
		"Or is it that we voted for this?",
		"Not with our rigged elections,",
	},
	{
		"but because we wanna be sedated.",
		"Because it's painful not to pretend,",
		"because we're cowards.",
		"- Elliot Alderson",
		"Mr. Robot",
	},
}

func producer(link chan string, wgp *sync.WaitGroup, i int) {
	defer wgp.Done()
	for _, v := range messages[i] {
		link <- v + " | producerID-" + strconv.Itoa(i)
	}
}

func consumer(link chan string, wgc *sync.WaitGroup, i int) {
	defer wgc.Done()
	for ele := range link {
		fmt.Printf("%d-consId-%s\n", i, ele)
	}
}

const consumerCount int = 3

func main(){
	link := make(chan string, 20)
	done := make(chan bool)

	wgp := new(sync.WaitGroup)
	wgc := new(sync.WaitGroup)

	go func(_link chan string, wgp *sync.WaitGroup) {
		defer close(_link)
		for i:=0; i<len(messages); i++ {
			wgp.Add(1)
			go producer(_link, wgp, i)
		}
		wgp.Wait()
	}(link, wgp)

	go func(_link chan string, wgc *sync.WaitGroup, done chan bool) {
		defer close(done)
		for i:=0; i<consumerCount; i++ {
			wgc.Add(1)
			go consumer(_link, wgc, i)
		}
		wgc.Wait()
		done <- true
	}(link, wgc, done)
	<- done
}
