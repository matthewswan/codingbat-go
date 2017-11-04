package main

import "fmt"

func main() {
	as := true
	bs := false
	monkeyTrouble(as, bs)
}

func monkeyTrouble(as, bs bool) bool {
	if as && bs || !as && !bs {
		fmt.Println("true")
		return true
	}
	fmt.Println("false")
	return false
}
