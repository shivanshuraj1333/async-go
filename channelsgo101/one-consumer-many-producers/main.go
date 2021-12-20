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

var consumers = 5

func produce(link chan string, i int, wg *sync.WaitGroup){
	defer wg.Done()
	for _, v := range messages[i] {
		link <- v + " | producerID-" + strconv.Itoa(i)
	}
}

func consume(link chan string, done chan bool) {
	defer close(done)
	for ele := range link {
		fmt.Println(ele)
	}
}

func main(){
	link := make(chan string, 18)
	done := make(chan bool)
	wg := new(sync.WaitGroup)
	go func(_link chan string, wg *sync.WaitGroup) {
		defer close(_link)
		for i:=0; i<len(messages); i++ {
			wg.Add(1)
			go produce(_link, i, wg)
		}
		wg.Wait()
	}(link, wg)
	go consume(link, done)
	<- done
}
