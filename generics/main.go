/*
Linked list implementation to understand generics in golang
*/
package main

import (
	"fmt"
)

// LinkedList node of Linked list
type LinkedList struct {
	value interface{}
	next  *LinkedList
}

// Len method returns lenght of the linked list
func (ll *LinkedList) Len() int {
	count := 0
	for node := ll; node != nil; node = node.next {
		count++
	}
	return count
}

// InsertAt method inserts new element/item at mentioned position in the linked list
func (ll *LinkedList) InsertAt(pos int, value interface{}) *LinkedList {
	if ll == nil || pos <= 0 {
		return &LinkedList{
			value: value,
			next:  ll,
		}
	}
	ll.next = ll.next.InsertAt(pos-1, value)
	return ll
}

// Append method appends element/item to the linked list
func (ll *LinkedList) Append(value interface{}) *LinkedList {
	return ll.InsertAt(ll.Len(), value)
}

// String method returns string representation of the linked list
func (ll *LinkedList) String() string {
	if ll == nil {
		return "nil"
	}
	return fmt.Sprintf("%v->%v", ll.value, ll.next.String())
}

func main() {
	fmt.Println("Generics:")
	var head *LinkedList
	head = head.Append("hello")
	head = head.Append(1)
	head = head.Append(1.2)
	fmt.Println(head.String())
	fmt.Println(head.Len())
}
