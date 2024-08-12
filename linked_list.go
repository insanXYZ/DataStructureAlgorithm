package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

type LinkedList struct {
	Node *Node
	Tail *Node
}

func (ll *LinkedList) add(i int) {
	newNode := &Node{
		Val:  i,
		Next: nil,
	}

	if ll.Node == nil {
		ll.Node = newNode
		ll.Tail = newNode
		return
	}

	ll.Tail.Next = newNode
	ll.Tail = newNode
}

func main() {
	ll := LinkedList{}
	ll.add(1)
	ll.add(2)
	ll.add(3)
	fmt.Println(ll.Node.Val, ll.Node.Next.Val, ll.Node.Next.Next.Val)
}
