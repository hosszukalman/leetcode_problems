package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

type TestCase struct {
	Input          string
	ExpectedResult int
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
		res := myAtoi(testCase.Input)
		if testCase.ExpectedResult != res {
			fmt.Printf("Something is not ok with the result: Test case: %#v, res: %+v\n", testCase, res)
		}
	}
}

// Accepted, Runtime: 0 ms, Memory Usage: 2.3 MB
// Runtime: 0 ms, faster than 100.00% of Go online submissions for String to Integer (atoi).
func myAtoi(s string) int {
	var chars []byte
	i := 0
	shouldNumber := false
	for i < len(s) {

		if shouldNumber && (s[i] < 48 || s[i] > 57) {
			break
		}

		switch s[i] {
		case 32: // space
			i++
			continue
		case 45: // "-""
			shouldNumber = true
			chars = append(chars, s[i])
			i++
			continue
		case 43: // "+"
			shouldNumber = true
			i++
			continue
		}

		if s[i] > 48 || s[i] < 57 {
			shouldNumber = true
			chars = append(chars, s[i])
		}

		i++
	}

	if len(chars) == 0 {
		return 0
	}

	res, _ := strconv.ParseInt(string(chars), 10, 32)

	return int(res)
}
