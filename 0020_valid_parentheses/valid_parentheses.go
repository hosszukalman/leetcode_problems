package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
		res := isValid(testCase.Input)
		if testCase.AcceptedResult != res {
			fmt.Printf("Something is not ok with the result: Test case: %#v, res: %+v\n", testCase, res)
		}
	}
}

// Accepted, Runtime: 0 ms, Memory Usage: 2.2 MB
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Valid Parentheses.
func isValid(s string) bool {
	parenthesesPairs := map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
	}
	var openParentheses []string
	for _, char := range s {
		charString := string(char)
		if charString == "(" || charString == "[" || charString == "{" {
			openParentheses = append(openParentheses, charString)
		} else {
			if len(openParentheses) == 0 {
				return false
			}
			lastOpenParenthes := openParentheses[len(openParentheses)-1]
			if parenthesesPairs[lastOpenParenthes] != charString {
				return false
			}
			openParentheses = openParentheses[:len(openParentheses)-1]
		}
	}
	return len(openParentheses) == 0
}
