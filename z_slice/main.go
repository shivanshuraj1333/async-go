package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := make([]int, 2)
	copy(b, a)
	b = append(b, 4)
	a[0] = 99
	fmt.Println("a:", a)
	fmt.Println("b:", b)
}
