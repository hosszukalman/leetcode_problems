package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"unicode"
)

type TestCase struct {
	Input          string
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
		res := detectCapitalUse(testCase.Input)
		if testCase.AcceptedResult != res {
			fmt.Printf("Something is not ok with the result: Input: %v, excepted: %v, res: %+v\n", testCase.Input, testCase.AcceptedResult, res)
		}
	}
}

// Accepted, Runtime: 0 ms, Memory Usage: 2.2 MB
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Detect Capital.
func detectCapitalUse(word string) bool {
	res := true
	lowerFound, upperFound := false, false
	for i := 1; i < len(word) && res; i++ {
		if unicode.IsLower(rune(word[i])) {
			lowerFound = true
		}
		if unicode.IsUpper(rune(word[i])) {
			upperFound = true
		}
		if unicode.IsLower(rune(word[0])) && upperFound {
			res = false
		}
		if unicode.IsUpper(rune(word[0])) && upperFound && lowerFound {
			res = false
		}
	}
	return res
}
