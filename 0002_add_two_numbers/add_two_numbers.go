package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (ln *ListNode) String() (s string) {
	if ln == nil {
		return "nil"
	}
	for ln != nil {
		s += fmt.Sprintf("%v ", ln.Val)
		if ln.Next != nil {
			s += "-> "
		}
		ln = ln.Next
	}

	return
}

type TestCase struct {
	List1          *ListNode
	List2          *ListNode
	ExpectedResult *ListNode
}

type TestCases []TestCase

func main() {
	var testCases TestCases
	jsonFile, err := ioutil.ReadFile("testCases.json")
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal([]byte(jsonFile), &testCases); err != nil {
		log.Fatal(err)
	}

	for _, testCase := range testCases {
		res := addTwoNumbers(testCase.List1, testCase.List2)
		if !reflect.DeepEqual(testCase.ExpectedResult, res) {
			fmt.Printf("Something is not ok with the result: list1: %v, list2: %v, excepted: %v, res: %+v\n", testCase.List1, testCase.List2, testCase.ExpectedResult, res)
		}
	}
}

// Accepted, Runtime: 8 ms, Memory Usage: 4.8 MB
// Runtime: 8 ms, faster than 85.88% of Go online submissions for Add Two Numbers.
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	type Result struct {
		*ListNode
		Head *ListNode
	}

	var calculateVal = func(v1, v2 int, plusOne bool) (val int, moreThanTen bool) {
		val = v1 + v2
		if plusOne {
			val++
		}
		if val >= 10 {
			val = val % 10
			moreThanTen = true
		}
		return
	}

	currentList1, currentList2 := l1, l2
	var res *Result = new(Result)

	moreThanTen := false
	for {
		var addToList *ListNode = new(ListNode)
		var v1, v2 int
		if currentList1 != nil {
			v1 = currentList1.Val
			currentList1 = currentList1.Next
		}
		if currentList2 != nil {
			v2 = currentList2.Val
			currentList2 = currentList2.Next
		}
		addToList.Val, moreThanTen = calculateVal(v1, v2, moreThanTen)

		if res.Head == nil {
			res.Head = addToList
			res.ListNode = addToList
		} else {
			currentNode := res.Head
			for currentNode.Next != nil {
				currentNode = currentNode.Next
			}
			currentNode.Next = addToList
		}
		if currentList1 == nil && currentList2 == nil && !moreThanTen {
			return res.ListNode
		}
	}
}

// Accepted, Runtime: 8 ms, Memory Usage: 4.8 MB
// Runtime: 8 ms, faster than 85.88% of Go online submissions for Add Two Numbers.
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	type Result struct {
		*ListNode
		Head *ListNode
	}

	currentList1, currentList2 := l1, l2
	var res *Result = new(Result)

	var carry int
	for {
		var addToList *ListNode = new(ListNode)
		var v1, v2 int
		if currentList1 != nil {
			v1 = currentList1.Val
			currentList1 = currentList1.Next
		}
		if currentList2 != nil {
			v2 = currentList2.Val
			currentList2 = currentList2.Next
		}
		addToList.Val, carry = (v1+v2+carry)%10, (v1+v2+carry)/10

		if res.Head == nil {
			if currentList1 == nil && currentList2 == nil && carry != 0 {
				addToList.Next = &ListNode{1, nil}
			}
			res.Head = addToList
			res.ListNode = addToList
		} else {
			currentNode := res.Head
			for currentNode.Next != nil {
				currentNode = currentNode.Next
			}
			if currentList1 == nil && currentList2 == nil && carry != 0 {
				addToList.Next = &ListNode{1, nil}
			}
			currentNode.Next = addToList
		}
		if currentList1 == nil && currentList2 == nil {
			return res.ListNode
		}
	}
}

// Accepted, Runtime: 12 ms, Memory Usage: 4.9 MB
// Runtime: 12 ms, faster than 56.00% of Go online submissions for Add Two Numbers.
func addTwoNumbers3(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	res := new(ListNode)
	sum := 0

	for current := res; l1 != nil || l2 != nil || sum != 0; current = current.Next {
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		current.Next = &ListNode{sum % 10, nil}
		sum /= 10
	}

	return res.Next
}
