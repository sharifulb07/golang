package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

func printList(head *Node) {
	current := head

	for current != nil {
		fmt.Println(current.Data)
		current=current.Next
	}

	fmt.Println("nil")

}

func main ()  {
	node3:=&Node{Data: 10, Next: nil }
	node2:=&Node{Data: 30, Next: node3}
	node1:=&Node{Data: 40, Next: node2}

	head:=node1

	printList(head)
}