package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
	"runtime"
	"sort"
	"sync"
)

type TestCase struct {
	Input          []int
	Target         int
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
		res := twoSum(testCase.Input, testCase.Target)
		if !reflect.DeepEqual(testCase.AcceptedResult, res) {
			fmt.Printf("Something is not ok with the result: Test case: %#v, res: %+v\n", testCase, res)
		}
	}

}

// Accepted, Runtime: 28 ms, Memory Usage: 3.7 MB
// faster than 34.13% of Go online submissions for Two Sum.
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
// faster than 24.51% of Go online submissions for Two Sum.
func twoSumCPUNumber(nums []int, target int) []int {

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

// Accepted, Runtime: 4 ms, Memory Usage: 4.2 MB
// faster than 97.00% of Go online submissions for Two Sum.
func twoSumOn(nums []int, target int) (res []int) {
	// [searchForValue]position
	hashMap := make(map[int]int, len(nums))

	for i := range nums {
		if position, found := hashMap[nums[i]]; found {
			return []int{position, i}
		} else {
			hashMap[target-nums[i]] = i
		}
	}
	return
}

// Accepted, Runtime: 8 ms, Memory Usage: 4.7 MB
// faster than 54.82% of Go online submissions for Two Sum.
func twoSumOnCPUNum(nums []int, target int) []int {
	n := runtime.NumCPU()
	resCh := make(chan []int)
	length := len(nums)

	incrementer := length / n
	odd := length % n

	// [searchForValue]position
	hashMap := make(map[int]int, len(nums))

	var lock = sync.RWMutex{}

	for procI := 1; procI <= n; procI++ {
		go func(procI int, nums []int, target int, length int, resCh chan []int) {
			start := (procI-1)*incrementer + odd + 1
			if procI == 1 {
				start = 0
			}

			for i := start; i < length; i++ {
				lock.Lock()
				if position, found := hashMap[nums[i]]; found {
					res := []int{position, i}
					resCh <- res
				} else {
					hashMap[target-nums[i]] = i
				}
				lock.Unlock()
			}

		}(procI, nums, target, length, resCh)
	}

	res := <-resCh

	close(resCh)
	return res
}

// Accepted, Runtime: 4 ms, Memory Usage: 3.9 MB
// faster than 97.00% of Go online submissions for Two Sum.
// Memory usage less than 70.04% of Go online submissions for Two Sum.
func twoSum(nums []int, target int) (res []int) {
	orderedNums := make([]int, len(nums))
	copy(orderedNums, nums)
	sort.Ints(orderedNums)

	a, b, start := 0, 0, 0
	end := len(orderedNums) - 1

	for {
		if start >= end {
			break
		}

		sum := orderedNums[start] + orderedNums[end]
		if sum < target {
			start++
		} else if sum > target {
			end--
		} else {
			a = orderedNums[start]
			b = orderedNums[end]
			break
		}
	}

	foundA, foundB := false, false
	for i, num := range nums {
		if num == a {
			res = append(res, i)
			foundA = true
			if a == b {
				continue
			}
		}
		if num == b {
			res = append(res, i)
			foundB = true
		}

		if foundA && foundB {
			return
		}
	}

	return
}
