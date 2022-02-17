package main

import "fmt"

// Node represents a node of linked list
type node struct {
	data int
	next *node
}

// LinkedList represents a linked list
type LinkedList struct {
	head *node
	len  int
}

func insert(head *node, data int)*node{
	n := &node{data: data}

	if head == nil {
		return n
	} else {
		n.next = head
		return n
	}

}

func PrintList(head *node) {
	for head != nil {
		fmt.Println(head.data, "->")
		head = head.next
	}
	fmt.Println(nil)
}

func main() {

	var link *node

	link  = insert(link, 1)
	link  = insert(link, 2)
	link  = insert(link, 3)


PrintList(link)

}
