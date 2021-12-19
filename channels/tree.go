package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int, 10), make(chan int, 10)

	go func(_ch1 chan int) {
		defer close(_ch1)
		Walk(t1, _ch1)
	}(ch1)

	go func(_ch2 chan int) {
		defer close(_ch2)
		Walk(t2, _ch2)
	}(ch2)

	for {
		val1, ok1 := <- ch1
		val2, ok2 := <- ch2

		if ok1 != ok2 || val1 != val2 {
			return false
		}
		if !ok1 && !ok2 {
			break
		}
	}
	return true
}

func main() {
	ch := make(chan int, 10)
	go func(_ch chan int){
		defer close(_ch)
		Walk(tree.New(1), _ch)
	}(ch)

	for {
		val, ok := <- ch
		if !ok {
			break
		} else {
			fmt.Println(val)
		}
	}

	fmt.Println(Same(tree.New(2), tree.New(2)))
	fmt.Println(Same(tree.New(2), tree.New(3)))
}