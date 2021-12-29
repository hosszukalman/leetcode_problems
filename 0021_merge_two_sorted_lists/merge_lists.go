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
		s += fmt.Sprintf("%v -> ", ln.Val)
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
		res := mergeTwoLists(testCase.List1, testCase.List2)
		if !reflect.DeepEqual(testCase.ExpectedResult, res) {
			fmt.Printf("Something is not ok with the result: list1: %v, list2: %v, excepted: %v, res: %+v\n", testCase.List1, testCase.List2, testCase.ExpectedResult, res)
		}
	}
}

// Accepted, Runtime: 0 ms, Memory Usage: 2.7 MB
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Merge Two Sorted Lists
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	type Result struct {
		*ListNode
		Head *ListNode
	}

	currentList1, currentList2 := list1, list2
	var res *Result = new(Result)

	for {
		var addToList *ListNode = new(ListNode)
		if currentList1 != nil {
			if currentList2 == nil {
				addToList.Val = currentList1.Val
				currentList1 = currentList1.Next
			} else if currentList1.Val <= currentList2.Val {
				addToList.Val = currentList1.Val
				currentList1 = currentList1.Next
			} else {
				addToList.Val = currentList2.Val
				currentList2 = currentList2.Next
			}
		} else if currentList2 != nil {
			addToList.Val = currentList2.Val
			currentList2 = currentList2.Next
		}
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
		if currentList1 == nil && currentList2 == nil {
			return res.ListNode
		}
	}
}
