package main

import "fmt"

func closeStrings(word1 string, word2 string) bool {
	// both words must have the same length
	// otherwise it is impossible to make
	// any of the operations
	if len(word1) != len(word2) {
		return false
	}

	m1 := make(map[byte]int, len(word1))
	m2 := make(map[byte]int, len(word2))

	for i := range word1 {
		m1[word1[i]]++
		m2[word2[i]]++
	}

	// both words must have the
	// same amount on unique characters
	if len(m1) != len(m2) {
		return false
	}

	n := make(map[int]int, len(m1))

	for k, c := range m1 {
		// all characters must be present
		// in both strings
		if _, ok := m2[k]; !ok {
			return false
		}

		n[c]++
	}

	for _, c := range m2 {
		n[c]--

		// if there are not enough characters
		// to transform then the strings are
		// not close
		if n[c] < 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(closeStrings("abc", "bca"))         // true
	fmt.Println(closeStrings("a", "aa"))            // false
	fmt.Println(closeStrings("cabbba", "abbccc"))   // true
	fmt.Println(closeStrings("cabbba", "aabbss"))   // false
	fmt.Println(closeStrings("uau", "ssx"))         // false
	fmt.Println(closeStrings("abbzzca", "babzzcz")) // false
}
