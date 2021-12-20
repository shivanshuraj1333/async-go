package main

import (
	"fmt"
)

var messages = []string{
	"The world itself's",
	"just one big hoax.",
	"Spamming each other with our",
	"running commentary of bullshit,",
	"masquerading as insight, our social media",
	"faking as intimacy.",
	"Or is it that we voted for this?",
	"Not with our rigged elections,",
	"but with our things, our property, our money.",
	"I'm not saying anything new.",
	"We all know why we do this,",
	"not because Hunger Games",
	"books make us happy,",
	"but because we wanna be sedated.",
	"Because it's painful not to pretend,",
	"because we're cowards.",
	"- Elliot Alderson",
	"Mr. Robot",
}

const consumerCount int = 3

func producer(link chan string) {
	defer close(link)
	for _, m := range messages {
		link <- m
	}
}

func consumer(link chan string, id int, done chan bool) {
	for ele := range link {
		fmt.Printf("%d - id: %s - msg\n",id, ele)
	}
	done <- true
}

func main() {
	link := make(chan string)
	done := make(chan bool, consumerCount)
	go producer(link)
	for i := 0; i<consumerCount; i++ {
		go consumer(link, i, done)
	}
	for i := 0; i<consumerCount; i++ {
		<- done
	}
}
