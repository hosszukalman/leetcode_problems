package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

type TestCase struct {
	Input          int
	ExpectedResult bool
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
		res := isPalindrome(testCase.Input)
		if testCase.ExpectedResult != res {
			fmt.Printf("Something is not ok with the result: Test case: %#v, res: %+v\n", testCase, res)
		}
	}
}

// Accepted, Runtime: 12 ms, Memory Usage: 5.2 MB
// faster than 79.98% of Go online submissions for Palindrome Number
func isPalindrome(x int) bool {
	intString := strconv.Itoa(x)
	len := len(intString)
	for i := 0; i < len/2; i++ {
		if intString[i] != intString[len-i-1] {
			return false
		}
	}

	return true
}
