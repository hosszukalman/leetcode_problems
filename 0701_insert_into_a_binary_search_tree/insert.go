package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type TestCase struct {
	InputTree      *TreeNode
	InputVal       int
	ExpectedResult *TreeNode
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
		res := insertIntoBST(testCase.InputTree, testCase.InputVal)
		if !reflect.DeepEqual(testCase.ExpectedResult, res) {
			printPreOrder(testCase.InputTree)
			fmt.Println()
			printPreOrder(testCase.ExpectedResult)
			fmt.Println()
			printPreOrder(res)
			fmt.Println()
			fmt.Println()
		}
	}
}

func printPreOrder(n *TreeNode) {
	if n == nil {
		return
	} else {
		fmt.Printf("%d ", n.Val)
		printPreOrder(n.Left)
		printPreOrder(n.Right)
	}
}

// Accepted, Runtime: 28 ms, Memory Usage: 7.5 MB
// Runtime: 28 ms, faster than 92.76% of Go online submissions for Insert into a Binary Search Tree.
func insert(root *TreeNode, val int) {
	if root.Val < val {
		if root.Right == nil {
			root.Right = &TreeNode{Val: val}
		} else {
			insert(root.Right, val)
		}
	} else {
		if root.Left == nil {
			root.Left = &TreeNode{Val: val}
		} else {
			insert(root.Left, val)
		}
	}
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	insert(root, val)

	return root
}

// Accepted, Runtime: 40 ms, Memory Usage: 7.8 MB
// Runtime: 40 ms, faster than 31.61% of Go online submissions for Insert into a Binary Search Tree.
func insertIntoBST2(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	} else {
		root.Right = insertIntoBST(root.Right, val)
	}

	return root
}
