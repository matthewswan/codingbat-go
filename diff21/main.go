package main

import (
	"fmt"
	"math"
)

func main() {
	n := 14
	diff21(n)
}

// if n if greater than 21, return double the absolute difference value
// else return the absolute difference value
func diff21(n int) float64 {
	if n > 21 {
		return doubleAbsDiff(n)
	}
	return absDiff(n)
}

func doubleAbsDiff(n int) float64 {
	fmt.Println(2 * math.Abs(float64(n)-21))
	return (2 * math.Abs(float64(n)-21))
}

func absDiff(n int) float64 {
	fmt.Println(math.Abs(float64(n) - 21))
	return (math.Abs(float64(n) - 21))

}
