/**
 * Package linkedlist
 * 单链表结构
 */
package linkedlist

import (
	"fmt"
	"reflect"
)

type ElementType interface{}

type INode struct {
	data ElementType
	next *INode
}

// LinkedList 单向链表 (头结点数据存放长度)
type LinkedList struct {
	Head *INode
}

// NewINode create a new node
func NewINode(data ElementType, next *INode) *INode {
	return &INode{data, next}
}

// NewLinkedList create a new LinkedList
func NewLinkedList() *LinkedList {
	head := &INode{0, nil}
	return &LinkedList{head}
}

// IsEmpty 判断是否为空链表
func (list *LinkedList) IsEmpty() bool {
	return list.Head.next == nil
}

// Length 遍历找长度
func (list *LinkedList) Length() int {
	return int(reflect.ValueOf((*list.Head).data).Int())
}

// Append 插入节点
func (list *LinkedList) Append(node *INode) {
	current := list.Head
	for {
		if current.next == nil {
			break
		}
		current = current.next
	}
	current.next = node
	list.sizeInc()
}

// Print 打印链表
func (list *LinkedList) Print() {
	if list.IsEmpty() {
		fmt.Println("empty list")
		return
	}
	current := list.Head.next
	fmt.Println("Print elements")
	i := 0
	for ; ; i++ {
		if current.next == nil {
			break
		}
		fmt.Printf("Node %d, value %v ", i, current.data)
		current = current.next
	}
	fmt.Printf("Node %d, value %v ", i+1, current.data)
	return
}

func (list *LinkedList) sizeInc() {
	len := int(reflect.ValueOf((*list.Head).data).Int())
	list.Head.data = len + 1
}
