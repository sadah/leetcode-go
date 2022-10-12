package code

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Print all nodes for singly-linked list.
func PrintAllListNode(node *ListNode) {
	fmt.Println("-------- PrintAllNode start")
	for node != nil {
		fmt.Printf("Node: %#v, %#v\n", node.Val, node.Next)
		node = node.Next
	}
	fmt.Println("-------- PrintAllNode end")
}

// Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

// Print all nodes for tree.
func PrintAllNode(root *Node) {
	fmt.Println("-------- PrintAllNode start")
	printNode(root)
	fmt.Println("-------- PrintAllNode end")
}

func printNode(node *Node) {
	fmt.Printf("Node: %#v, %#v\n", node.Val, node.Children)
	for _, n := range node.Children {
		printNode(n)
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
