package code

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func PrintAllNode(node *ListNode) {
	fmt.Println("PrintAllNode start")
	for node != nil {
		fmt.Printf("Node: %#v, %#v\n", node.Val, node.Next)
		node = node.Next
	}
	fmt.Println("PrintAllNode end")
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
