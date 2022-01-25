package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type TestCase struct {
	Input          []int
	AcceptedResult bool
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
		res := validMountainArray(testCase.Input)
		if testCase.AcceptedResult != res {
			fmt.Printf("Something is not ok with the result: Input: %v, excepted: %v, res: %+v\n", testCase.Input, testCase.AcceptedResult, res)
		}
	}
}

// Accepted, Runtime: 20 ms, Memory Usage: 6.7 MB
// Runtime: 20 ms, faster than 98.44% of Go online submissions for Valid Mountain Array.
func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}

	if arr[0] > arr[1] {
		return false
	}
	if arr[len(arr)-1] > arr[len(arr)-2] {
		return false
	}

	res, shouldDecrease := true, false
	for i := 1; i < len(arr) && res; i++ {
		if arr[i] == arr[i-1] {
			res = false
		}

		if arr[i] < arr[i-1] {
			shouldDecrease = true
		}

		if arr[i] > arr[i-1] && shouldDecrease {
			res = false
		}
	}

	return res
}
