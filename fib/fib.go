package fib

import (
	"fmt"
	"math"
)

// RetrieveNthFibonacci receives a number between 0 - 1474
// as a posith in the Fibonacci's Sequence and
// returns a String representation for the value at this position
// This function has O(1) performance thanks to the application
// of PHI formula for Fibanacci.
//
// References that I followed:
// https://www.quora.com/What-is-the-complexity-of-an-algorithm-that-calculates-the-nth-number-of-the-Fibonacci-series
// https://www.quora.com/How-can-we-find-the-nth-term-in-a-Fibonacci-sequence-What-is-the-sum-of-n-terms-in-that-series/answer/Naveen-Kbs
// https://en.wikipedia.org/wiki/Fibonacci_number
// https://ideone.com/fM6htW
// http://www.maths.surrey.ac.uk/hosted-sites/R.Knott/Fibonacci/fibtable.html#:~:text=The%20Fibonacci%20series,-is%20formed%20by&text=0%2B1%3D1%20so%20the,Fibonacci%20Calculator%2C%20written%20in%20JavaScript.
// http://mersennus.net/fibonacci/fibonacci.txt
// https://stackoverflow.com/questions/45023864/ruby-floatdomainerror-infinity-when-calculating-fibonacci-number
func RetrieveNthFibonacci(position uint16) (string, error) {
	rootFive := math.Sqrt(5)
	PHI := (1 + rootFive) / 2

	if position < uint16(3) {
		return "1", nil
	}

	if position > uint16(1474) {
		return "", fmt.Errorf("ERROR: invalid Fibonacci Sequence index position [%v]", position)
	}

	power := math.Pow(PHI, float64(position))
	baseValue := (power / rootFive) + 0.5
	baseValueFloor := math.Floor(baseValue)
	result := fmt.Sprintf("%0.f", baseValueFloor)

	return fmt.Sprintf("%v", result), nil
}
