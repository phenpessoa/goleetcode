package main

import "fmt"

var vowels = []rune{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}

func isVowel(r rune) bool {
	for _, v := range vowels {
		if v == r {
			return true
		}
	}
	return false
}

func halvesAreAlike(s string) bool {
	a := s[:len(s)/2]
	b := s[len(s)/2:]

	var countA, countB int
	for _, r := range a {
		if isVowel(r) {
			countA++
		}
	}

	for _, r := range b {
		if isVowel(r) {
			countB++
		}
	}

	return countA == countB
}

func main() {
	fmt.Println(halvesAreAlike("book"))
}
