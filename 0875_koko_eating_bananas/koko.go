package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"sort"
)

type TestCase struct {
	InputPiles     []int
	InputHours     int
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
		res := minEatingSpeed(testCase.InputPiles, testCase.InputHours)
		if testCase.AcceptedResult != res {
			fmt.Printf("Something is not ok with the result: InputPiles: %v, InputHours: %v, excepted: %v, res: %+v\n", testCase.InputPiles, testCase.InputHours, testCase.AcceptedResult, res)
		}
	}
}

// Accepted, Runtime: 39 ms, Memory Usage: 7.6 MB
// Runtime: 39 ms, faster than 78.57% of Go online submissions for Koko Eating Bananas.
func minEatingSpeed1(piles []int, h int) int {
	sort.Ints(piles)
	p1, p2 := 1, piles[len(piles)-1]
	mid := (p1 + p2) / 2

	for p1 <= p2 {
		hNeed := 0
		mid = (p1 + p2) / 2
		for _, val := range piles {
			hNeed += int(math.Ceil(float64(val) / float64(mid)))
			if hNeed > h {
				break
			}
		}
		if hNeed > h {
			p1 = mid + 1
		} else {
			p2 = mid - 1
		}
	}

	return p1
}

// Accepted, Runtime: 21 ms, Memory Usage: 7.1 MB
// Runtime: 21 ms, faster than 97.14% of Go online submissions for Koko Eating Bananas.
func minEatingSpeed(piles []int, h int) int {
	p1, p2 := 1, 1000000000
	mid := (p1 + p2) / 2

	for p1 <= p2 {
		hNeed := 0
		mid = (p1 + p2) / 2
		for _, val := range piles {
			hNeed += int(math.Ceil(float64(val) / float64(mid)))
			if hNeed > h {
				break
			}
		}
		if hNeed > h {
			p1 = mid + 1
		} else {
			p2 = mid - 1
		}
	}

	return p1
}

func minEatingSpeed3(piles []int, h int) int {
	p1, p2 := 1, 1000000000
	mid := (p1 + p2) >> 1

	for p1 <= p2 {
		hNeed := 0
		mid = (p1 + p2) >> 1
		for _, val := range piles {
			hNeed += int(math.Ceil(float64(val) / float64(mid)))
			if hNeed > h {
				break
			}
		}
		if hNeed > h {
			p1 = mid + 1
		} else {
			p2 = mid - 1
		}
	}

	return p1
}
