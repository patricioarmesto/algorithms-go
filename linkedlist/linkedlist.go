package main

import "fmt"

type Node struct {
	key  interface{}
	next *Node
}

type List struct {
	head *Node
}

func (l *List) Insert(key interface{}) {

	node := &Node{key: key}
	if l.head == nil {
		l.head = node
	} else {
		l := l.head
		for l.next != nil {
			l = l.next
		}

		l.next = node
	}
}

func (l *List) Show() {
	list := l.head
	for list != nil {
		fmt.Printf("%+v -> ", list.key)
		list = list.next
	}

	fmt.Println("nil")
	fmt.Println()
}

func main() {

	list := List{}
	list.Insert(1)
	list.Insert(2)
	list.Show()
}
