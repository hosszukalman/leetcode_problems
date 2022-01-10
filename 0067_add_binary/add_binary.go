package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"strconv"
)

type TestCase struct {
	Input1         string
	Input2         string
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
		res := addBinary(testCase.Input1, testCase.Input2)
		if testCase.ExpectedResult != res {
			fmt.Printf("Something is not ok with the result: Test case: %#v, res: %+v\n", testCase, res)
		}
	}
}

// Accepted, Runtime: 0 ms, Memory Usage: 2.3 MB
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Add Binary.
func addBinary(a string, b string) string {
	n1 := new(big.Int)
	n1.SetString(a, 2)
	n2 := new(big.Int)
	n2.SetString(b, 2)
	sum := new(big.Int)
	return sum.Add(n1, n2).Text(2)
}

type Runes []rune

func (str Runes) ReverseString() (res Runes) {
	l := len(str)
	res = make(Runes, l)
	for i := 0; i <= l/2; i++ {
		res[i], res[l-1-i] = str[l-1-i], str[i]
	}
	return res
}

// Accepted, Runtime: 4 ms, Memory Usage: 2.8 MB
// Runtime: 4 ms, faster than 20.86% of Go online submissions for Add Binary.
func addBinary2(a string, b string) string {
	var res string
	i, j, carry := len(a)-1, len(b)-1, 0
	for i >= 0 || j >= 0 {
		sum := carry
		var val int64
		if i >= 0 {
			val, _ = strconv.ParseInt(string(a[i]), 2, 2)
			sum += int(val)
			i--
		}
		if j >= 0 {
			val, _ = strconv.ParseInt(string(b[j]), 2, 2)
			sum += int(val)
			j--
		}

		if sum > 1 {
			carry = 1
		} else {
			carry = 0
		}

		res += fmt.Sprintf("%v", sum%2)
	}

	if carry != 0 {
		res += "1"
	}

	return string(Runes(res).ReverseString())
}
