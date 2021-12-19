package main

import "golang.org/x/tour/tree"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	defer close(ch)
	var wait func(t *tree.Tree, ch chan int)
	wait = func(t *tree.Tree, ch chan int) {
		if t != nil {
			Walk(t.Left)
			ch <- t.Value
			Walk(t.Right)
		}
	}
	wait(t, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
}

func main() {
}

