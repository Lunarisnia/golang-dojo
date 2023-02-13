package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	defer close(ch)
	var inOrder func(t *tree.Tree)
	inOrder = func(t *tree.Tree) {
		if t == nil {
			return
		}

		inOrder(t.Left)
		ch <- t.Value
		inOrder(t.Right)
	}
	inOrder(t)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	// Run Logic
	ch, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch)
	go Walk(t2, ch2)
	for {
		x, ok1 := <-ch
		y, ok2 := <-ch2
		if x != y {
			return false
		}
		if !ok1 || !ok2 {
			break
		}
	}
	return true
}

func main() {
	c := make(chan int)
	go Walk(tree.New(1), c)
	fmt.Println(<-c)
	for v := range c {
		fmt.Println(v)
	}
	fmt.Println(Same(tree.New(1), tree.New(1)))
}
