package main

import "fmt"

func main() {
	root := CreateTree(2)
	root.Left = CreateTree(5)
	root.Right = CreateTree(6)
	root.Left.Left = CreateTree(9)
	root.Left.Right = CreateTree(90)
	root.Right.Left = CreateTree(24)
	root.Right.Right = CreateTree(32)
	PreOrder(root)
	InOrder(root)
	PostOrder(root)
}

type Tree struct {
	Value int
	Left  *Tree
	Right *Tree
}

func CreateTree(value int) *Tree {
	return &Tree{value, nil, nil}
}

//前序遍历
func PreOrder(tree *Tree) {
	if tree != nil {
		fmt.Println(tree.Value)
		PreOrder(tree.Left)
		PreOrder(tree.Right)
	}
}

//中序遍历
func InOrder(tree *Tree) {
	if tree != nil {

		PreOrder(tree.Left)
		fmt.Print(" ", tree.Value)
		PreOrder(tree.Right)
	}
}

//后序遍历
func PostOrder(tree *Tree) {
	if tree != nil {
		PreOrder(tree.Left)
		PreOrder(tree.Right)
		fmt.Print(" ", tree.Value)
	}
}
