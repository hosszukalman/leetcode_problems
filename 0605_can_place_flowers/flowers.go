package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type TestCase struct {
	InputFlowerbed []int
	InputNumber    int
	AcceptedResult bool
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
		res := canPlaceFlowers(testCase.InputFlowerbed, testCase.InputNumber)
		if testCase.AcceptedResult != res {
			fmt.Printf("Something is not ok with the result: InputFlowerbed: %v, inputNumber: %v, excepted: %v, res: %+v\n", testCase.InputFlowerbed, testCase.InputNumber, testCase.AcceptedResult, res)
		}
	}
}

// Accepted, Runtime: 16 ms, Memory Usage: 6.1 MB
// Runtime: 16 ms, faster than 82.14% of Go online submissions for Can Place Flowers.
func canPlaceFlowers1(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}

	if n == 1 && len(flowerbed) == 1 && flowerbed[0] == 0 {
		return true
	}
	freeSlots, lastFlowerPosition := 0, -1

	i := 0
	for i < len(flowerbed) {
		if flowerbed[i] != 0 {
			lastFlowerPosition = i
			i++
			continue
		}

		if i == 1 && lastFlowerPosition == -1 {
			freeSlots++
			lastFlowerPosition = 0
		}

		if i == len(flowerbed)-1 && i-lastFlowerPosition == 2 {
			freeSlots++
			lastFlowerPosition = i
		}

		if i-lastFlowerPosition == 3 {
			freeSlots++
			lastFlowerPosition = i - 1
		}

		if freeSlots == n {
			return true
		}
		i++
	}

	return false
}

// Accepted, Runtime: 16 ms, Memory Usage: 6.1 MB
// Runtime: 16 ms, faster than 82.14% of Go online submissions for Can Place Flowers.
func canPlaceFlowers2(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}

	if n == 1 && len(flowerbed) == 1 && flowerbed[0] == 0 {
		return true
	}
	freeSlots, lastFlowerPosition := 0, -1

	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] != 0 {
			lastFlowerPosition = i
			continue
		}

		if i == 1 && lastFlowerPosition == -1 {
			freeSlots++
			lastFlowerPosition = 0
		}

		if i == len(flowerbed)-1 && i-lastFlowerPosition == 2 {
			freeSlots++
			lastFlowerPosition = i
		}

		if i-lastFlowerPosition == 3 {
			freeSlots++
			lastFlowerPosition = i - 1
		}

		if freeSlots == n {
			return true
		}
	}

	return false
}

// Accepted, Runtime: 16 ms, Memory Usage: 6.1 MB
// Runtime: 16 ms, faster than 82.14% of Go online submissions for Can Place Flowers.
func canPlaceFlowers(flowerbed []int, n int) bool {
	for i := 0; i < len(flowerbed) && n > 0; i++ {
		if flowerbed[i] != 0 {
			continue
		}
		prev, next := 0, 0
		if i > 0 {
			prev = flowerbed[i-1]
		}
		if i < len(flowerbed)-1 {
			next = flowerbed[i+1]
		}

		if prev == 0 && next == 0 {
			flowerbed[i] = 1
			n--
		}
	}

	return n == 0
}

// Accepted, Runtime: 16 ms, Memory Usage: 6.1 MB
// Runtime: 16 ms, faster than 82.14% of Go online submissions for Can Place Flowers.
func canPlaceFlowers4(flowerbed []int, n int) bool {
	zeros, ans := 1, 0

	for _, val := range flowerbed {
		if val == 0 {
			zeros++
		} else {
			ans += (zeros - 1) / 2
			zeros = 0
		}

		if ans >= n {
			return true
		}
	}
	return ans+(zeros/2) >= n
}
