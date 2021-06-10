package main

import "fmt"

type TreeNode struct {
	Val       int
	LeftNode  *TreeNode
	RightNode *TreeNode
}

func main() {
	node := &TreeNode{5, nil, nil}
	tree := &TreeNode{1, &TreeNode{2, node, nil}, &TreeNode{9, nil, nil}}
	newTree := invertTree(tree)
	fmt.Println(newTree)
	count1 := count(tree)
	fmt.Println(count1)
}

// 前序遍历二叉树  /* 二叉树遍历框架 */
var preOrder, inOrder, postOrder []int

func traverse(root *TreeNode) *TreeNode {

	//前序遍历 根左右
	traverse(root.LeftNode)
	//中序遍历 左根右
	traverse(root.RightNode)
	//后序遍历 左右根
	return root
}

// 定义：count(root) 返回以 root 为根的树有多少节点
func count(root *TreeNode) int {
	// base case
	if root == nil {
		return 0
	}
	// 自己加上子树的节点数就是整棵树的节点数
	return 1 + count(root.LeftNode) + count(root.RightNode)
}

// 翻转二叉树
//我们发现只要把二叉树上的每一个节点的左右子节点进行交换，最后的结果就是完全翻转之后的二叉树。
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	tmp := root.LeftNode
	root.LeftNode = root.RightNode
	root.RightNode = tmp

	//让左右节点继续翻转他们的子节点
	invertTree(root.LeftNode)
	invertTree(root.RightNode)

	return root
}
