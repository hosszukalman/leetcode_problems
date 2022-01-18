package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type TestCase struct {
	InputHaystack  string
	InputNeedle    string
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
		res := strStr(testCase.InputHaystack, testCase.InputNeedle)
		if testCase.AcceptedResult != res {
			fmt.Printf("Something is not ok with the result: Test case: %#v, res: %+v\n", testCase, res)
		}
	}
}

// Accepted, Runtime: 320 ms, Memory Usage: 2.6 MB
// Runtime: 320 ms, faster than 21.07% of Go online submissions for Implement strStr().
func strStr1(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	for i := range haystack {
		if haystack[i] == needle[0] {
			if i+len(needle) > len(haystack) {
				return -1
			}
			found := true
			for j := range needle {
				if haystack[i+j] != needle[j] {
					found = false
					break
				}
			}
			if found {
				return i
			}
		}
	}

	return -1
}

// Accepted, Runtime: 0 ms, Memory Usage: 2.5 MB
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Implement strStr().
func strStr2(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	if len(needle) > len(haystack) {
		return -1
	}

	return strings.Index(haystack, needle)
}

// Accepted, Runtime: 4 ms, Memory Usage: 2.5 MB
// Runtime: 8 ms, faster than 81.94% of Go online submissions for Implement strStr().
func strStr(haystack string, needle string) int {
	needleLength := len(needle)

	if needleLength == 0 {
		return 0
	}

	if needleLength > len(haystack) {
		return -1
	}

	for i := 0; i <= len(haystack)-needleLength; i++ {
		if haystack[i:i+needleLength] == needle {
			return i
		}
	}

	return -1
}
