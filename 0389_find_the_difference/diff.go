package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"sync"
)

type TestCase struct {
	Input1         string
	Input2         string
	AcceptedResult string
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
		res := findTheDifference(testCase.Input1, testCase.Input2)
		b := []byte(testCase.AcceptedResult)
		if b[0] != res {
			fmt.Printf("Something is not ok with the result: Test case: %+v, res: %v\n", testCase, res)
		}
	}
}

// Accepted, Runtime: 0 ms, Memory Usage: 2.1 MB
// Runtime: 0 ms, faster than 100% of Go online submissions for Find the Difference.
func findTheDifferenceWithMap(s string, t string) (res byte) {
	originalChars, newChars := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if _, found := originalChars[s[i]]; found {
			originalChars[s[i]]++
		} else {
			originalChars[s[i]] = 1
		}
	}
	for i := 0; i < len(t); i++ {
		if _, found := newChars[t[i]]; found {
			newChars[t[i]]++
		} else {
			newChars[t[i]] = 1
		}
	}

	for char, newCount := range newChars {
		if originalCount, found := originalChars[char]; found {
			if newCount != originalCount {
				return char
			}
		} else {
			return char
		}
	}

	return
}

// Accepted, Runtime: 0 ms, Memory Usage: 2.6 MB
// Runtime: 0 ms, faster than 100% of Go online submissions for Find the Difference.
func findTheDifference(s string, t string) (res byte) {
	originalChars, newChars := make(map[byte]int), make(map[byte]int)

	stringToMap := func(str string) map[byte]int {
		m := make(map[byte]int)
		for i := 0; i < len(str); i++ {
			if _, found := m[str[i]]; found {
				m[str[i]]++
			} else {
				m[str[i]] = 1
			}
		}
		return m
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		originalChars = stringToMap(s)
	}()
	go func() {
		defer wg.Done()
		newChars = stringToMap(t)
	}()

	wg.Wait()

	for char, newCount := range newChars {
		if originalCount, found := originalChars[char]; found {
			if newCount != originalCount {
				return char
			}
		} else {
			return char
		}
	}

	return
}

// Accepted, Runtime: 0 ms, Memory Usage: 2.2 MB
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Find the Difference.
func findTheDifferenceXor(s string, t string) (res byte) {
	for i := 0; i < len(s); i++ {
		res ^= s[i]
	}
	for i := 0; i < len(t); i++ {
		res ^= t[i]
	}
	return
}

// Accepted, Runtime: 0 ms, Memory Usage: 2.1 MB
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Find the Difference.
func findTheDifferenceXorConcurrency(s string, t string) (res byte) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < len(s); i++ {
			res ^= s[i]
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < len(t); i++ {
			res ^= t[i]
		}
	}()
	wg.Wait()
	return
}
