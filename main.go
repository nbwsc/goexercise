package main

import (
	"fmt"
	"nbwsc/goexercise/linkedlist"
)

func main() {
	list := linkedlist.NewLinkedList()
	for i := 0; i < 10; i++ {
		node := linkedlist.NewINode(fmt.Sprintf("str%d", i), nil)
		list.Append(node)
	}
	list.Print()
	fmt.Println(list.Length())
}
