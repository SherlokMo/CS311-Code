package main

import "fmt"

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
