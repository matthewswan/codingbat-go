package main

import "fmt"

func main() {
	t := true
	h := 22
	parrotTrouble(t, h)
}

func parrotTrouble(t bool, h int) bool {
	if (h < 7 || h > 20) && t == true {
		fmt.Println("true")
		return true
	}
	fmt.Println("false")
	return false
}
