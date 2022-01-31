package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type TestCase struct {
	Input          [][]int
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
		res := maximumWealth(testCase.Input)
		if testCase.AcceptedResult != res {
			fmt.Printf("Something is not ok with the result: TestCase: %+v, res: %v\n", testCase, res)
		}
	}
}

// Accepted, Runtime: 3 ms, Memory Usage: 3.1 MB
// Runtime: 3 ms, faster than 93.03% of Go online submissions for Richest Customer Wealth.
func maximumWealth1(accounts [][]int) (max int) {
	for _, account := range accounts {
		accountWealth := 0
		for _, val := range account {
			accountWealth += val
		}
		if accountWealth > max {
			max = accountWealth
		}
	}
	return
}

// Accepted, Runtime: 0 ms, Memory Usage: 3.3 MB
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Richest Customer Wealth.
func maximumWealth(accounts [][]int) (max int) {
	calculateWealth := func(account []int, sumCh chan int) {
		sum := 0
		for _, val := range account {
			sum += val
		}
		sumCh <- sum
	}

	for _, account := range accounts {
		sumCh := make(chan int)
		go calculateWealth(account, sumCh)
		accountWealth := <-sumCh
		if accountWealth > max {
			max = accountWealth
		}
		close(sumCh)
	}
	return
}
