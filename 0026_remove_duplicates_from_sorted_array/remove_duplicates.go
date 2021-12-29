package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
)

type TestCase struct {
	Input          []int
	ExpectedResult struct {
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
		var input = make([]int, len(testCase.Input))
		copy(input, testCase.Input)
		res := removeDuplicates(input)
		if testCase.ExpectedResult.Result != res || !reflect.DeepEqual(input, testCase.ExpectedResult.ModifiedInput) {
			fmt.Printf("Something is not ok with the result: Test case: %+v, result: %+v, modified input: %+v\n", testCase, res, input)
		}
	}
}

// Accepted, Runtime: 118 ms, Memory Usage: 4.6 MB
// Runtime: 118 ms, faster than 5.08% of Go online submissions for Remove Duplicates from Sorted Array.
func removeDuplicates1(nums []int) (res int) {
	prev := -101
	for i := 0; i < len(nums) && nums[i] != -101; i++ {
		if prev != nums[i] {
			res++
			prev = nums[i]
		} else {
			nums = append(nums[:i], nums[i+1:]...)
			nums = append(nums, -101)
			i--
		}
	}
	return
}

// Accepted, Runtime: 56 ms, Memory Usage: 4.4 MB
// Runtime: 56 ms, faster than 13.48% of Go online submissions for Remove Duplicates from Sorted Array.
func removeDuplicates2(nums []int) (res int) {
	prev := -101
	for i := 0; i < len(nums); i++ {
		if prev != nums[i] {
			res++
			prev = nums[i]
		} else {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}
	return
}

// Accepted, Runtime: 5 ms, Memory Usage: 4.6 MB
// Runtime: 5 ms, faster than 86.46% of Go online submissions for Remove Duplicates from Sorted Array.
func removeDuplicates(nums []int) (res int) {
	prev := -101
	for i := 0; i < len(nums); i++ {
		if prev != nums[i] {
			res++
			prev = nums[i]
		} else {
			j := 1
			for j = 1; i+j < len(nums); j++ {
				if prev != nums[i+j] {
					break
				}
			}
			nums = append(nums[:i], nums[i+j:]...)
			i--
		}
	}
	return
}
