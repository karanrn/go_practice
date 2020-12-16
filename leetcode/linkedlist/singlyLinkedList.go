package main

import (
	"fmt"
	"time"
)

type listNode struct {
	Val  int
	Next *listNode
}

func addTwoNumbers(l1 *listNode, l2 *listNode) *listNode {
	var sum, carry, rem int
	var res *listNode
	for {
		var temp *listNode
		if l1 != nil && l2 != nil {
			sum = l1.Val + l2.Val + carry
			carry = (sum / 10) % 10
			rem = (sum) % 10
			temp = &listNode{Val: rem}

			l1 = l1.Next
			l2 = l2.Next
		} else if l1 != nil && l2 == nil {
			sum = l1.Val + carry
			carry = (sum / 10) % 10
			rem = (sum) % 10
			temp = &listNode{Val: rem}

			l1 = l1.Next
			l2 = nil
		} else if l1 == nil && l2 != nil {
			sum = l2.Val + carry
			carry = (sum / 10) % 10
			rem = (sum) % 10
			temp = &listNode{Val: rem}

			l1 = nil
			l2 = l2.Next
		}
		if res == nil {
			res = temp
		} else {
			temp.Next = res
			res = temp
			/*
				current := res
				for current.Next != nil {
					current = current.Next
				}
				current.Next = temp
			*/
		}

		if l1 == nil && l2 == nil {
			if carry != 0 {
				temp = &listNode{Val: carry}
				temp.Next = res
				res = temp
				/*
					current := res
					for current.Next != nil {
						current = current.Next
					}
					current.Next = temp
				*/
			}
			break
		}
	}

	return res
}

func main() {
	node1 := listNode{Val: 9}
	node2 := listNode{Val: 9}
	node3 := listNode{Val: 9}
	node1.Next = &node2
	node2.Next = &node3

	node4 := node2
	start := time.Now()
	res := addTwoNumbers(&node1, &node4)
	elapsed := time.Now().Sub(start)

	fmt.Printf("Program took: %d ns\n", elapsed.Nanoseconds())

	if res != nil {
		for {
			fmt.Println(res.Val)
			if res.Next != nil {
				res = res.Next
			} else {
				break
			}
		}
	}

}
