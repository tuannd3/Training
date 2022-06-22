package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int
	fmt.Scanf("%d", &a)
	if checkPalindrome(a) {
		fmt.Println(a, "is a Palindrome")
	} else {
		fmt.Println(a, "not is a Palindrome")
	}
}

func checkPalindrome(a int) bool {
	strA := strconv.Itoa(a)
	for i := 0; i < len(strA)/2; i++ {
		if strA[i] != strA[len(strA)-1-i] {
			return false
		}
	}
	return true
}
