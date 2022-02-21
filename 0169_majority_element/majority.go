package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
)

type TestCase struct {
	Input          []int
	AcceptedResult int
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
		res := majorityElement(testCase.Input)
		if testCase.AcceptedResult != res {
			fmt.Printf("Something is not ok with the result: Test case: %+v, res: %+v\n", testCase, res)
		}
	}
}

// Accepted, Runtime: 12 ms, Memory Usage: 6.1 MB
// Runtime: 12 ms, faster than 98.68% of Go online submissions for Majority Element.
func majorityElement1(nums []int) int {
	store := make(map[int]int)

	for _, n := range nums {
		if _, found := store[n]; found {
			store[n]++
		} else {
			store[n] = 1
		}
		if store[n] > len(nums)/2 {
			return n
		}
	}
	return 0
}

// Accepted, Runtime: 16 ms, Memory Usage: 6 MB
// Runtime: 16 ms, faster than 89.72% of Go online submissions for Majority Element.
func majorityElement2(nums []int) int {
	sort.Ints(nums)
	minCount := len(nums) / 2

	i := 0
	for {
		if nums[i] == nums[i+minCount] {
			return nums[i]
		} else {
			for j := i; j <= i+minCount; j++ {
				if nums[i] != nums[j] {
					i = j
					break
				}
			}
		}
	}
}

// Accepted, Runtime: 16 ms, Memory Usage: 6 MB
// Runtime: 16 ms, faster than 89.72% of Go online submissions for Majority Element.
func majorityElement3(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

// Accepted, Runtime: 12 ms, Memory Usage: 6.1 MB
// Runtime: 12 ms, faster than 98.68% of Go online submissions for Majority Element.
func majorityElement(nums []int) int {
	store := make(map[int]int)

	for _, n := range nums {
		store[n]++
		if store[n] > len(nums)/2 {
			return n
		}
	}
	return -1
}
