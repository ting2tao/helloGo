/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 删除第n个节点  相当于遍历删除 len-n+1位置的节点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	listNodeLength := 0
	cur := &ListNode{0, head}
	tmp := cur
	for ; head != nil; head = head.Next {
		listNodeLength++
	}
	for i := 0; i < listNodeLength-n; i++ {
		tmp = tmp.Next
	}
	tmp.Next = tmp.Next.Next
	return cur.Next
}

func main() {
	head := &ListNode{1, &ListNode{
		2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}},
	}}
	list := removeNthFromEnd(head, 2)
	var arr []int
	for ; list != nil; list = list.Next {
		arr = append(arr, list.Val)
	}
	fmt.Println(arr)
}
