package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
		res := romanToInt(testCase.Input)
		if testCase.ExpectedResult != res {
			fmt.Printf("Something is not ok with the result: Test case: %#v, res: %+v\n", testCase, res)
		}
	}
}

// Accepted, Runtime: 16 ms, Memory Usage: 6.4 MB
// faster than 13.64% of Go online submissions for Roman to Integer
func romanToInt1(s string) (res int) {
	type CharMap map[string]int
	charMap := CharMap{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	specialChars := map[string]CharMap{
		"I": {
			"V": 4,
			"X": 9,
		},
		"X": {
			"L": 40,
			"C": 90,
		},
		"C": {
			"D": 400,
			"M": 900,
		},
	}
	for i := 0; i < len(s); i++ {
		if specialCharDef, specialCharfound := specialChars[string(s[i])]; specialCharfound && i < len(s)-1 {
			if value, specialCharfound := specialCharDef[string(s[i+1])]; specialCharfound {
				res += value
				i++
				continue
			} else {
				res += charMap[string(s[i])]
			}
		} else {
			res += charMap[string(s[i])]
		}
	}
	return
}

// Accepted, Runtime: 4 ms, Memory Usage: 3 MB
// faster than 88.21% of Go online submissions for Roman to Integer.
func romanToInt(s string) (res int) {
	type CharMap map[string]int
	charMap := CharMap{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	for i := 0; i < len(s); i++ {
		if string(s[i]) == string("I") && i < len(s)-1 {
			if string(s[i+1]) == string("V") {
				res += 4
				i++
				continue
			} else if string(s[i+1]) == string("X") {
				res += 9
				i++
				continue
			} else {
				res += charMap[string(s[i])]
			}
		} else if string(s[i]) == string("X") && i < len(s)-1 {
			if string(s[i+1]) == string("L") {
				res += 40
				i++
				continue
			} else if string(s[i+1]) == string("C") {
				res += 90
				i++
				continue
			} else {
				res += charMap[string(s[i])]
			}
		} else if string(s[i]) == string("C") && i < len(s)-1 {
			if string(s[i+1]) == string("D") {
				res += 400
				i++
				continue
			} else if string(s[i+1]) == string("M") {
				res += 900
				i++
				continue
			} else {
				res += charMap[string(s[i])]
			}
		} else {
			res += charMap[string(s[i])]
		}
	}
	return
}
