package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type TestCase struct {
	Input          int
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
		res := addDigits(testCase.Input)
		if testCase.AcceptedResult != res {
			fmt.Printf("Something is not ok with the result: Test case: %+v, res: %+v\n", testCase, res)
		}
	}
}

// Accepted, Runtime: 0 ms, Memory Usage: 2.2 MB
// Runtime: 0 ms, faster than 100% of Go online submissions for Add Digits.
func addDigits1(num int) (res int) {
	res = num

	splitDigits := func(n int) (res []int) {
		tail := n % 10
		res = append(res, tail)
		for div := n / 10; div != 0; div /= 10 {
			tail = div % 10
			res = append(res, tail)
		}

		return
	}

	sum := func(nums []int) (res int) {
		for _, val := range nums {
			res += val
		}

		return
	}

	for digits := splitDigits(num); len(digits) > 1; {
		res = sum(digits)
		digits = splitDigits(res)
	}

	return
}

// Accepted, Runtime: 0 ms, Memory Usage: 2.1 MB
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Add Digits.
func addDigits2(num int) int {
	return 1 + (num-1)%9
}

// Accepted, Runtime: 3 ms, Memory Usage: 3.1 MB
// Runtime: 3 ms, faster than 48.61% of Go online submissions for Add Digits.
func addDigits(num int) int {
	splitDigits := func(n int, digitCh chan<- int, lenCh chan<- int, doneCh chan<- bool) {
		length := 1
		tail := n % 10
		digitCh <- tail
		for div := n / 10; div != 0; div /= 10 {
			tail = div % 10
			length++
			digitCh <- tail
		}
		lenCh <- length
		doneCh <- true
	}

	sumFunc := func(digitCh <-chan int, sumCh chan<- int, doneCh <-chan bool) {
		n := 0
		for {
			select {
			case <-doneCh:
				sumCh <- n
				return
			case add := <-digitCh:
				n += add
			}
		}
	}

	digitCh, lenCh, sumCh := make(chan int), make(chan int), make(chan int)
	doneCh := make(chan bool)

	sum := num
	for {
		go splitDigits(sum, digitCh, lenCh, doneCh)
		go sumFunc(digitCh, sumCh, doneCh)
		length := <-lenCh
		sum = <-sumCh
		if length == 1 {
			break
		}

	}

	close(digitCh)
	close(lenCh)
	close(sumCh)
	close(doneCh)

	return sum
}
