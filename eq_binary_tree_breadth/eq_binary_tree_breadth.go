package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	queue := make([]*tree.Tree, 0)
	queue = append(queue, t)
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		ch <- current.Value

		if current.Left != nil {
			queue = append(queue, current.Left)
		}
		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	m := make(map[int]int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for v := range ch1 {
		m[v] = m[v] + 1
	}
	for v := range ch2 {
		m[v] = m[v] - 1
	}
	for v := range m {
		if m[v] > 0 {
			return false
		}
	}
	return true
}

func main() {

	ch := make(chan int)
	go Walk(tree.New(1), ch)
	fmt.Print("Walk:")
	for v := range ch {
		fmt.Print(" ", v)
	}
	fmt.Println()

	fmt.Println("Same Expect true:", Same(tree.New(1), tree.New(1)))
	fmt.Println("Same Expect false:", Same(tree.New(1), tree.New(2)))

}
