package main

import "fmt"

/**
	N is the length of string st

	Algorithm isPalindrom(st, n):
		l <- 0
		r <- n - 1
		while l < r
			if st[l] != st[r]:
				return false
			r <- r - 1
			l <- l + 1
		endwhile

		return true
**/

func palindromeIterative(st string) bool {

	for l, r := 0, len(st)-1; l < r; {
		if st[l] != st[r] {
			return false
		}
		l++
		r--
	}

	return true
}

func main() {

	s := "ROTOR"
	fmt.Println(palindromeIterative(s))

}
