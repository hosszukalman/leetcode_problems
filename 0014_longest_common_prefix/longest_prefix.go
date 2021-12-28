package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
)

type TestCase struct {
	Input          []string
	ExpectedResult string
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
		res := longestCommonPrefix(testCase.Input)
		if testCase.ExpectedResult != res {
			fmt.Printf("Something is not ok with the result: Test case: %#v, res: %+v\n", testCase, res)
		}
	}
}

// Accepted, Runtime: 0 ms, Memory Usage: 2.4 MB
// faster than 100.00% of Go online submissions for Longest Common Prefix
func longestCommonPrefix(strs []string) (res string) {
	if len(strs) == 0 {
		return
	}
	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) < len(strs[j])
	})
	n := len(strs[0])
	for j := 0; j < n; j++ {
		var prev byte
		for i := range strs {
			if i != 0 && strs[i][j] != prev {
				return
			}

			prev = strs[i][j]
		}

		res += string(prev)
	}
	return
}
