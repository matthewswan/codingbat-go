package main

import "fmt"

func main() {
	w := true
	v := false
	sleepIn(w, v)
}

func sleepIn(w, v bool) bool {
	if !w || v {
		fmt.Println("true")
		return true
	}
	fmt.Println("false")
	return false
}
