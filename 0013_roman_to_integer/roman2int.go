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
