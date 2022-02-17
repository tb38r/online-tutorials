package main

import "fmt"

//define node
type node struct {
	next *node
	data int
}

//define linked list
type linkedList struct {
	head   *node
	length int
}

//Func that allow adding nodes to the list
func (l *linkedList) prepend(value int) {

	newNode := node{data: value}

	if l.head != nil { //adding to a pre-existing list or not?
		newNode.next = l.head //l.head holds a reference to the first node in a given list 
		l.head = &newNode
		l.length++
	} else {
		l.head = &newNode
		l.length++
	}


}

func (l *linkedList) PrintList() {
	if l.head == nil { //no nodes in list, return as there's nothing to print
		return //fmt.Println(no nodes to print)
	}
	currentNode := l.head
	for currentNode != nil {
		fmt.Println(currentNode.data)
		currentNode = currentNode.next
	}

}

func (l *linkedList) deleteWithValue(value int) {
	if l.length == 0 {
		return
	}

if l.head.data == value{
	l.head = l.head.next
	l.length--
	return
}

previousToDelete := l.head
for previousToDelete.next.data != value{
	if previousToDelete.next.next == nil{
		return
	}
	previousToDelete=previousToDelete.next
}
previousToDelete.next = previousToDelete.next.next
l.length--

}
