package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	b := []int{1, 2, 3, 9}
	result(a, b)
}

func result(a []int, b []int) {
	m := make(map[int]bool)
	if len(a) < len(b) {
		for _, i := range a {
			m[i] = true
		}
		for _, j := range b {
			if m[j] == true {
				fmt.Println(j)
			}
		}
	} else {
		for _, i := range b {
			m[i] = true
		}
		for _, j := range a {
			if m[j] == true {
				fmt.Println(j)
			}
		}
	}
}
