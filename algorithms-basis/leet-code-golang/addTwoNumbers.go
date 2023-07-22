package main
// import (  
//     "fmt"
// )

 type ListNode struct {
    Val  int
    Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    linkedList := &ListNode{}
	current := linkedList
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		val1, val2 := 0, 0

		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}

		total := val1 + val2 + carry
		carry = total / 10
		digits := total % 10

		current.Next = &ListNode{Val: digits}
		current = current.Next

	}
	return linkedList.Next
}

// func main() {
// 	l1 := &ListNode{Val: 2}
// 	l1.Next = &ListNode{Val: 4}
// 	l1.Next.Next = &ListNode{Val: 3}

// 	l2 := &ListNode{Val: 5}
// 	l2.Next = &ListNode{Val: 6}
// 	l2.Next.Next = &ListNode{Val: 4}

// 	result := addTwoNumbers(l1, l2)

// 	for result != nil {
// 		fmt.Printf("%d ", result.Val)
// 		result = result.Next
// 	}
// }