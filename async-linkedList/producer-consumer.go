package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Node struct {
	val int
	next *Node
}

type List struct {
	head *Node
	len int
}

func sendRequest(v int, link chan int, wg *sync.WaitGroup) {
	// mark routine completion
	defer wg.Done()
	// add value
	link <- v
	// random processing time between 1-1000secs
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000-1) + 1) )
}

func receiveRequest(db *List, link chan int) {
	fmt.Println("Adding nodes in order: ")
	for ele := range link {
		addNode(db, ele)
	}
}

func addNode(db *List, ele int) {
	var m sync.Mutex
	m.Lock()
	if db.head == nil {
		db.head = &Node{ele, nil}
		db.len++
		fmt.Printf("Head is at Node %d \n", db.head.val)
		fmt.Printf("Adding nodes:\n %d ", db.head.val)
	} else {
		c := db.head
		for {
			if c.next == nil {
				c.next = &Node{ele, nil}
				fmt.Printf("%d ", c.next.val)
				db.len++
				break
			}
			c = c.next
		}
	}
	m.Unlock()
}

// iterating over linked list and printing values
func (db *List) printLinkedList() {
	c := db.head
	fmt.Printf("Length of linked list db is %d \n", db.len)

	id := 0
	for {
		if c == nil {
			break
		} else {
			fmt.Printf("[id: %d | val: %v] -> ", id, c.val)
			id++
			c = c.next
		}
	}
}

func main(){
	// test ids to process individual id by sendRequest function
	id := []int{0, 1, 2, 3, 4, 5, 6, 7}
	// channel between sender and receiver
	link := make(chan int, 16)
	// wg to track sendRequest
	wgSendDone := new(sync.WaitGroup)
	// channel to mark receiving process as done
	done := make(chan bool)
	// database having linkedList as node for every processed id
	db := List{nil, 0}

	// initiating a sender go routine to send individual requests
	go func(_link chan int){
		defer close(_link)
		for _, v := range id {
			wgSendDone.Add(1)
			go sendRequest(v, _link, wgSendDone)
		}
		// waiting for completion of all send requests
		wgSendDone.Wait()
	}(link)

	// initiating a reception go routine to receive all processed id result in linkedList
	go func(_db List, _link chan int, _done chan bool){
		defer close(_done)
		receiveRequest(&db, link)
		done <- true
	}(db, link, done)

	// waiting till all id results are received
	<- done

	fmt.Println("\n\n--- --- Done receiving! --- ---")
	fmt.Println("\nPrinting linkedList values")

	// printing node values of db linkedList
	db.printLinkedList()

	// *** end of program ***
}
/*
  Output:

	Adding nodes in order:
	Head is at Node 7
	Adding nodes:
	 7 3 4 5 6 0 2 1

	--- --- Done receiving! --- ---

	Printing linkedList values
	Length of linked list db is 8
	[id: 0 | val: 7] -> [id: 1 | val: 3] -> [id: 2 | val: 4] -> [id: 3 | val: 5] -> [id: 4 | val: 6] -> [id: 5 | val: 0] -> [id: 6 | val: 2] -> [id: 7 | val: 1] ->

*/