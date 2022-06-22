package main

import "fmt"

func main() {
	for i := 2; i <= 100000; i++ {
		if checkSNT(i) {
			fmt.Println(i)
		}
	}
}

func checkSNT(a int) bool {
	if a < 2 {
		return false
	}
	for i := 2; i*i <= a; i++ {
		if a%i == 0 {
			return false
		}
	}
	return true
}
