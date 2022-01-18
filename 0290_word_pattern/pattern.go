package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type TestCase struct {
	InputPattern   string
	InputString    string
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
		res := wordPattern(testCase.InputPattern, testCase.InputString)
		if testCase.AcceptedResult != res {
			fmt.Printf("Something is not ok with the result: Input pattern: %v, input string: %v, excepted: %v, res: %+v\n", testCase.InputPattern, testCase.InputString, testCase.AcceptedResult, res)
		}
	}
}

// Accepted, Runtime: 0 ms, Memory Usage: 2 MB
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Word Pattern.
func wordPattern(pattern string, s string) bool {
	sSlice := strings.Split(s, " ")
	if len(pattern) != len(sSlice) {
		return false
	}

	var foundStrings []string

	hashMap := make(map[rune]string)
	for i, char := range pattern {
		if val, found := hashMap[char]; found {
			if sSlice[i] != val {
				return false
			}
		} else {
			for j := range foundStrings {
				if foundStrings[j] == sSlice[i] {
					return false
				}
			}
			hashMap[char] = sSlice[i]
			foundStrings = append(foundStrings, sSlice[i])
		}
	}

	return true
}
