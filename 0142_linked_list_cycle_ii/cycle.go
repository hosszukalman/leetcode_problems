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

func printNode(ln *ListNode) {
	if ln == nil {
		fmt.Print("nil")
	}
	for ln != nil {
		fmt.Printf("%v -> ", ln.Val)
		ln = ln.Next
	}
}

type TestCase struct {
	Input          *ListNode
	AcceptedResult *ListNode
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
		res := detectCycle(testCase.Input)
		if !reflect.DeepEqual(testCase.AcceptedResult, res) {
			fmt.Printf("Something is not ok with the result: input: %v, excepted: %v, res: %+v\n", testCase.Input, testCase.AcceptedResult, res)
		}
	}
}

// Accepted, Runtime: 8 ms, Memory Usage: 6.2MB
// Runtime: 8 ms, faster than 50.28% of Go online submissions for Linked List Cycle II.
func detectCycle(head *ListNode) (res *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	hashMap := make(map[*ListNode]int)
	for head != nil {
		if _, found := hashMap[head.Next]; found {
			return head.Next
		}
		hashMap[head] = head.Val
		head = head.Next
	}

	return
}
