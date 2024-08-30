package anagram

import (
	"reflect"
	"sort"
	"testing"
)

func TestAnagrams(t *testing.T) {
	for _, test := range testCases {
		actual := Anagrams(test.word, test.candidates)
		sort.Strings(actual)
		sort.Strings(test.expected)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("Anagrams(%q, %v) = %v, want %v", test.word, test.candidates, actual, test.expected)
		}
	}
}

var testCases = []struct {
	word       string
	candidates []string
	expected   []string
}{
	{
		word:       "diaper",
		candidates: []string{"hello", "world", "zombies", "pants"},
		expected:   []string{},
	},
	{
		word:       "ant",
		candidates: []string{"tan", "stand", "at"},
		expected:   []string{"tan"},
	},
	{
		word:       "galea",
		candidates: []string{"eagle"},
		expected:   []string{},
	},
	{
		word:       "good",
		candidates: []string{"dog", "goody"},
		expected:   []string{},
	},
	{
		word:       "listen",
		candidates: []string{"enlists", "google", "inlets", "banana"},
		expected:   []string{"inlets"},
	},
	{
		word:       "allergy",
		candidates: []string{"gallery", "ballerina", "regally", "clergy", "largely", "leading"},
		expected:   []string{"gallery", "regally", "largely"},
	},
	{
		word:       "Orchestra",
		candidates: []string{"cashregister", "Carthorse", "radishes"},
		expected:   []string{"Carthorse"},
	},
	{
		word:       "ΑΒΓ",
		candidates: []string{"ΒΓΑ", "ΒΓΔ", "γβα"},
		expected:   []string{"ΒΓΑ", "γβα"},
	},
	{
		word:       "ΑΒΓ",
		candidates: []string{"ABΓ"},
		expected:   []string{},
	},
}
