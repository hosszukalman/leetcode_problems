package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
)

type TestCase struct {
	InputNums      []int
	InputVal       int
	AcceptedResult struct {
		Result        int
		ModifiedInput []int
	}
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
		var input = make([]int, len(testCase.InputNums))
		copy(input, testCase.InputNums)
		res := removeElement(input, testCase.InputVal)
		if testCase.AcceptedResult.Result != res {
			fmt.Printf("Something is not ok with the result: Test case: %+v, result: %+v, modified input: %+v\n", testCase, res, input)
		} else {
			ok := true
			sort.Ints(input)
			for i := 0; i < res; i++ {
				if input[i] != testCase.AcceptedResult.ModifiedInput[i] {
					ok = false
					break
				}
			}
			if !ok {
				fmt.Printf("Something is not ok with the result: Test case: %+v, result: %+v, modified input: %+v\n", testCase, res, input)
			}
		}
	}
}

// Accepted, Runtime: 0 ms, Memory Usage: 2.2 MB
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Remove Element.
func removeElement1(nums []int, val int) (res int) {
	const marker = 101
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums[i] = marker
		} else {
			res++
		}
	}
	sort.Ints(nums)
	return
}

// Accepted, Runtime: 0 ms, Memory Usage: 2.2 MB
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Remove Element.
func removeElement(nums []int, val int) (res int) {
	i := 0
	for i < len(nums) {
		if nums[i] == val {
			if i == len(nums)-1 {
				nums = nums[:i]
			} else {
				nums = append(nums[:i], nums[i+1:]...)
			}
			continue
		} else {
			res++
		}
		i++
	}
	return
}
