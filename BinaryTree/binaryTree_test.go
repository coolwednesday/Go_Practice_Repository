package BinaryTree

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/tour/tree"
	"testing"
)

func TestBinaryTree(t *testing.T) {
	type numberOfNodes struct {
		tree1 int
		tree2 int
	}
	tests := []struct {
		name   string
		input  numberOfNodes
		output bool
	}{
		{name: "For Null Trees", input: numberOfNodes{0, 0}, output: true},
		{name: "For Trees with equal value sequence", input: numberOfNodes{1, 1}, output: true},
		{name: "For Trees with unequal value sequence", input: numberOfNodes{1, 2}, output: false},
	}
	//tree.New(k) function creates a tree with 10 nodes with values k, 2k.......10k
	for _, test := range tests {
		assert.Equal(t, test.output, same(tree.New(test.input.tree1), tree.New(test.input.tree2)))
	}
}
