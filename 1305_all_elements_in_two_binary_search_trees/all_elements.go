package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
	"sort"
	"sync"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type TestCase struct {
	Input1         *TreeNode
	Input2         *TreeNode
	AcceptedResult []int
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
		res := getAllElements(testCase.Input1, testCase.Input2)
		if !reflect.DeepEqual(testCase.AcceptedResult, res) {
			fmt.Printf("Something is not ok with the result: Input1: %v, Input1: %v, excepted: %v, res: %+v\n", testCase.Input1, testCase.Input2, testCase.AcceptedResult, res)
		}
	}
}

func preOrder1(tn *TreeNode, vals *[]int) {
	if tn != nil {
		*vals = append(*vals, tn.Val)
		if tn.Left != nil || tn.Right != nil {
			preOrder1(tn.Left, vals)
			preOrder1(tn.Right, vals)
		}
	}
}

// Accepted, Runtime: 217 ms, Memory Usage: 9.6 MB
// Runtime: 217 ms, faster than 17.39% of Go online submissions for All Elements in Two Binary Search Trees.
func getAllElements1(root1 *TreeNode, root2 *TreeNode) (res []int) {
	preOrder1(root1, &res)
	preOrder1(root2, &res)
	sort.Ints(res)
	return
}

func preOrder2(tn *TreeNode) (res []int) {
	if tn != nil {
		if tn.Left != nil || tn.Right != nil {
			res = append(res, preOrder2(tn.Left)...)
			res = append(res, preOrder2(tn.Right)...)
		}

		res = append(res, tn.Val)
	}

	return
}

// Accepted, Runtime: 375 ms, Memory Usage: 37.3 MB
// Runtime: 375 ms, faster than 5.80% of Go online submissions for All Elements in Two Binary Search Trees.
func getAllElements2(root1 *TreeNode, root2 *TreeNode) (res []int) {
	res = preOrder2(root1)
	res = append(res, preOrder2(root2)...)
	sort.Ints(res)
	return
}

func preOrder3(tn *TreeNode, vals *wrapper) {
	if tn != nil {
		vals.Append(tn.Val)
		if tn.Left != nil || tn.Right != nil {
			preOrder3(tn.Left, vals)
			preOrder3(tn.Right, vals)
		}
	}
}

type wrapper struct {
	mu   sync.RWMutex
	vals []int
}

func (w *wrapper) Append(val int) {
	w.mu.Lock()
	w.vals = append(w.vals, val)
	w.mu.Unlock()
}

// Accepted, Runtime: 112 ms, Memory Usage: 8.1 MB
// Runtime: 112 ms, faster than 60.87% of Go online submissions for All Elements in Two Binary Search Trees.
func getAllElements3(root1 *TreeNode, root2 *TreeNode) []int {
	var wg sync.WaitGroup
	res := wrapper{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		preOrder3(root1, &res)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		preOrder3(root2, &res)
	}()
	wg.Wait()
	sort.Ints(res.vals)
	return res.vals
}

func preOrder4(tn *TreeNode, vals *[]int) {
	if tn != nil {
		*vals = append(*vals, tn.Val)
		if tn.Left != nil || tn.Right != nil {
			preOrder4(tn.Left, vals)
			preOrder4(tn.Right, vals)
		}
	}
}

// Accepted, Runtime: 104 ms, Memory Usage: 8.3 MB
// Runtime: 104 ms, faster than 84.06% of Go online submissions for All Elements in Two Binary Search Trees.
func getAllElements4(root1 *TreeNode, root2 *TreeNode) []int {
	var wg sync.WaitGroup
	var list1, list2 []int

	wg.Add(1)
	go func() {
		defer wg.Done()
		preOrder4(root1, &list1)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		preOrder4(root2, &list2)
	}()
	wg.Wait()

	sort.Ints(list1)
	sort.Ints(list2)

	var res []int
	i, j := 0, 0
	for i < len(list1) || j < len(list2) {
		if j >= len(list2) || i < len(list1) && list1[i] <= list2[j] {
			res = append(res, list1[i])
			i++
		} else if i >= len(list1) || j < len(list2) && list1[i] > list2[j] {
			res = append(res, list2[j])
			j++
		}
	}
	return res
}

func preOrder(tn *TreeNode, vals *[]int) {
	if tn != nil {
		*vals = append(*vals, tn.Val)
		if tn.Left != nil || tn.Right != nil {
			preOrder(tn.Left, vals)
			preOrder(tn.Right, vals)
		}
	}
}

func merge(list1, list2 []int) (res []int) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		sort.Ints(list1)
	}()
	go func() {
		defer wg.Done()
		sort.Ints(list2)
	}()
	wg.Wait()

	i, j := 0, 0
	for i < len(list1) || j < len(list2) {
		if j >= len(list2) || i < len(list1) && list1[i] <= list2[j] {
			res = append(res, list1[i])
			i++
		} else if i >= len(list1) || j < len(list2) && list1[i] > list2[j] {
			res = append(res, list2[j])
			j++
		}
	}

	return
}

// Accepted, Runtime: 92 ms, Memory Usage: 8 MB
// Runtime: 92 ms, faster than 98.46% of Go online submissions for All Elements in Two Binary Search Trees.
func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	var wg sync.WaitGroup
	var list1, list2 []int

	wg.Add(2)
	go func() {
		defer wg.Done()
		preOrder(root1, &list1)
	}()
	go func() {
		defer wg.Done()
		preOrder(root2, &list2)
	}()
	wg.Wait()

	return merge(list1, list2)
}
