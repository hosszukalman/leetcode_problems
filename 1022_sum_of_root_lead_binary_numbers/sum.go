package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// func (tn *TreeNode) String() (s string) {
// 	if tn == nil {
// 		// fmt.Printf("nill", tn.Val)
// 	} else {
// 		// fmt.Printf("%d -> ", tn.Val)
// 		// s += fmt.Sprintf("%d -> ", tn.Val)
// 		// tn.Left.String()
// 		// tn.Right.String()
// 	}
// 	// for tn.Left != nil || tn.Right != nil {
// 	// 	s += fmt.Sprintf("%v -> ", ln.Val)
// 	// 	ln = ln.Next
// 	// }

// 	return
// }

// func (ln *ListNode) String() (s string) {
// 	if ln == nil {
// 		return "nil"
// 	}
// 	for ln != nil {
// 		s += fmt.Sprintf("%v -> ", ln.Val)
// 		ln = ln.Next
// 	}

// 	return
// }

type TestCase struct {
	Input          *TreeNode
	ExpectedResult int
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
		res := sumRootToLeaf(testCase.Input)
		if testCase.ExpectedResult != res {
			fmt.Printf("Something is not ok with the result: Input: %v, excepted: %v, res: %+v\n", testCase.Input, testCase.ExpectedResult, res)
		}
	}
}

// cannot define new methods on non-local type leetcode.TreeNode (solution.go)
func (tn *TreeNode) PreOrder(s string, f func(s string)) {
	if tn != nil {
		s += strconv.Itoa(tn.Val)
		if tn.Left == nil && tn.Right == nil {
			f(s)
		} else {
			tn.Left.PreOrder(s, f)
			tn.Right.PreOrder(s, f)
		}
	}
}

func sumRootToLeaf1(root *TreeNode) (res int) {
	root.PreOrder("", func(s string) {
		i, _ := strconv.ParseInt(s, 2, 32)
		res += int(i)
	})
	return
}

func preOrder(tn *TreeNode, s string, f func(s string)) {
	if tn != nil {
		s += strconv.Itoa(tn.Val)
		if tn.Left == nil && tn.Right == nil {
			f(s)
		} else {
			preOrder(tn.Left, s, f)
			preOrder(tn.Right, s, f)
		}
	}
}

// Accepted, Runtime: 0 ms, Memory Usage: 3.4 MB
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Sum of Root To Leaf Binary Numbers.
func sumRootToLeaf(root *TreeNode) (res int) {
	preOrder(root, "", func(s string) {
		i, _ := strconv.ParseInt(s, 2, 32)
		res += int(i)
	})
	return
}
