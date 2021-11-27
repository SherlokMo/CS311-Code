package main

import "fmt"

func palindromeRecursive(st string, l, r int) bool {
	if l == r {
		return true
	}

	if st[l] == st[r] {
		return palindromeRecursive(st, l+1, r-1)
	}

	return false

}

func main() {

	s := "ROTOR"
	fmt.Println(palindromeRecursive(s, 0, len(s)-1))

}
