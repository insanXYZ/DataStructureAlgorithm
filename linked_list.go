package main

import (
	"fmt"
)

type Node struct {
	Val  int
	Next *Node
}

type LinkedList struct {
	Node *Node
}

func (ll *LinkedList) add(i int) {
	newNode := &Node{
		Val:  i,
		Next: nil,
	}

	if ll.Node == nil {
		ll.Node = newNode
		return
	}

	current := ll.Node
	for current != nil {
		if current.Next == nil {
			break
		}
		current = current.Next
	}
	current.Next = newNode
}

func (ll *LinkedList) remove(i int) {
	current := ll.Node
	if current.Val == i {
		ll.Node = current.Next
		return
	}

	for current != nil {
		if current.Next.Val == i {
			current.Next = current.Next.Next
			return
		}
		current = current.Next
	}
}

func main() {
	ll := LinkedList{}
	ll.add(1)
	ll.add(2)
	ll.add(3)
	fmt.Println(ll.Node.Val, ll.Node.Next.Val, ll.Node.Next.Next.Val)
	ll.remove(2)
	fmt.Println(ll.Node, ll.Node.Next, ll.Node.Next.Next)
}
