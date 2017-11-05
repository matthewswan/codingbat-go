package main

import "fmt"

func main() {
	n1 := 5
	n2 := 5
	if n1 == n2 {
		sumDouble(n1, n2)
	} else {
		sum(n1, n2)
	}
}

func sumDouble(n1, n2 int) int {
	fmt.Println(2 * (n1 + n2))
	return 2 * (n1 + n2)
}

func sum(n1, n2 int) int {
	fmt.Println(n1 + n2)
	return n1 + n2
}
