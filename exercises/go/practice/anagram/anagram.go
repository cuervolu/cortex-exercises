package anagram

import (
	"sort"
	"strings"
)

// Anagrams finds all anagrams of a given word from a list of candidates
func Anagrams(word string, candidates []string) []string {
	// TODO: Implement the anagram detection logic
	return nil
}

// Helper function to sort characters in a string
func sortString(s string) string {
	chars := strings.Split(strings.ToLower(s), "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}
