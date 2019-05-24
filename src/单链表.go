package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode)append(val int){
	for l.Next != nil{

	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	newLN := &ListNode{}
	carry := 0
	p, q, current := l1, l2, newLN
	for {
		if p == nil && q == nil{
			if carry > 0 && current.Next != nil{
				current = current.Next
				current.Val = carry
			}else{
				current.Next = nil
			}
			break
		}
		if current.Next != nil{
			current = current.Next
		}
		var x,y int
		if p == nil {x = 0}else {x = p.Val}
		if q == nil {y = 0}else {y = q.Val}
		currVal := x + y + carry
		if currVal >= 10{
			carry = 1
			currVal = currVal - 10

		}else{
			carry = 0
		}
		if p != nil{p = p.Next}else {p = nil}
		if q != nil{q = q.Next}else {q = nil}
		current.Val = currVal
		current.Next = &ListNode{}
	}
	return newLN
}

func addDGTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var curr *ListNode = nil
	l1Res := sum(l1, 0, 1)
	l2Res := sum(l2, 0, 1)
	res := l1Res + l2Res
	for _, str := range strconv.Itoa(res){
		va, _ := strconv.Atoi(string(str))
		curr = &ListNode{Val:va, Next:curr}
	}
	return curr
}

func sum(l1 *ListNode, res int, wei int) int{
	var x int
	if l1 == nil {x = 0} else {x = l1.Val}
	res = res + x * wei
	wei = wei * 10
	if l1 == nil || l1.Next == nil{
		return res
	}else {
		return sum(l1.Next, res, wei)
	}
}

func main() {
	var a1 = []int{1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1}
	var a2 = []int{5,6,4}
	var l1, l2 *ListNode = nil, nil
	fmt.Println(a2)
	for _, i := range a1{
		//if index == len(a1) - 1{
		//
		//}
		l1.Val = i
		l1.Next = &ListNode{}
	}
	for _, i := range a2{
		l2 = &ListNode{Val:i, Next:l2}
	}
	res := addDGTwoNumbers(l1, l2)
	fmt.Println(res)
}
