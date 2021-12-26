package main

import (
	"fmt"
	"reflect"
	"runtime"
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
func twoSum0n2(nums []int, target int) (res []int) {
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

// Accepted, Runtime: 32 ms, Memory Usage: 3.7 MB
func twoSum(nums []int, target int) []int {

	n := runtime.NumCPU()
	resCh := make(chan []int)
	length := len(nums)

	incrementer := length / n
	odd := length % n

	for procI := 1; procI <= n; procI++ {
		go func(procI int, nums []int, target int, length int, resCh chan []int) {
			start := (procI-1)*incrementer + odd + 1
			if procI == 1 {
				start = 0
			}

			for i := start; i < length; i++ {
				for j := i + 1; j < length; j++ {
					if nums[i]+nums[j] == target {
						res := []int{i, j}
						resCh <- res
					}
				}
			}

		}(procI, nums, target, length, resCh)
	}

	res := <-resCh

	close(resCh)
	return res
}
