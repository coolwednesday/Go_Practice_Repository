package BinaryTree

import "golang.org/x/tour/tree"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
// inorder traversal and sending the node value in the channel
func walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	if t.Left != nil {
		walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	go walk(t1, ch1)
	go walk(t2, ch2)
	//Comparing values from both the channels for the same sequence
	for i := 0; i < 10; i++ {
		x := <-ch1
		y := <-ch2
		if x != y {
			return false
		}
	}
	return true
}
