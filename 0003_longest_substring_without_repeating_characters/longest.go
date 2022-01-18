package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type TestCase struct {
	Input          string
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
		res := lengthOfLongestSubstring(testCase.Input)
		if testCase.AcceptedResult != res {
			fmt.Printf("Something is not ok with the result: Test case: %#v, res: %+v\n", testCase, res)
		}
	}
}

// Accepted, Runtime: 46 ms, Memory Usage: 3.1 MB
// Runtime: 46 ms, faster than 24.24% of Go online submissions for Longest Substring Without Repeating Characters.
func lengthOfLongestSubstring1(s string) int {
	if s == "" {
		return 0
	}
	hashMap := make(map[rune]int)
	longest := 0
	for i, char := range s {
		if position, found := hashMap[char]; !found {
			hashMap[char] = i
		} else {
			for storedChar, storedPosition := range hashMap {
				if storedPosition <= position {
					delete(hashMap, storedChar)
				}
			}
			hashMap[char] = i
		}
		if len(hashMap) > longest {
			longest = len(hashMap)
		}
	}
	return longest
}

// Accepted, Runtime: 0 ms, Memory Usage: 3.2 MB
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Longest Substring Without Repeating Characters.
func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	hashMap := make(map[rune]int)
	longest, start := 0, 0
	for i, char := range s {
		if position, found := hashMap[char]; found && position >= start {
			start = position + 1
		}
		if i-start+1 > longest {
			longest = i - start + 1
		}
		hashMap[char] = i
	}
	return longest
}
