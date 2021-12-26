package main

import (
	"fmt"
	"reflect"
)

type TestCase struct {
	Input          []int
	Target         int
	ExceptedResult []int
}

type TestCases []TestCase

func main() {
	testCases := TestCases{
		TestCase{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		TestCase{[]int{3, 2, 4}, 6, []int{1, 2}},
		TestCase{[]int{3, 3}, 6, []int{0, 1}},
	}

	for _, testCase := range testCases {
		res := twoSum(testCase.Input, testCase.Target)
		if !reflect.DeepEqual(testCase.ExceptedResult, res) {
			fmt.Printf("Something is not ok with the result: Input: %#v, res: %+v\n", testCase.Input, res)
		}
	}

}

// Accepted, Runtime: 28 ms, Memory Usage: 3.7 MB
func twoSum(nums []int, target int) (res []int) {
out:
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				res = []int{i, j}
				break out
			}
		}
	}
	return
}
