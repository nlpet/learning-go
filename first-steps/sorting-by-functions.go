package main

import (
	"fmt"
	"sort"
	"strings"
)

const vowels string = "aeiou"

type ByLength []string
type ByNumVowels []string

// By length
func (s ByLength) Len() int           { return len(s) }
func (s ByLength) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s ByLength) Less(i, j int) bool { return len(s[i]) < len(s[j]) }

// By number of vowels
func (s ByNumVowels) Len() int      { return len(s) }
func (s ByNumVowels) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s ByNumVowels) Less(i, j int) bool {
	return numVowels(s[i]) > numVowels(s[j])
}

func numVowels(s string) int {
	numVowels := 0

	for _, r := range s {
		if strings.Contains(vowels, string(r)) {
			numVowels++
		}
	}
	return numVowels
}

func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(ByLength(fruits))
	fmt.Println(fruits)

	words := []string{"aaaaaa", "aa", "bb", "bba", "bbaa", "bbaaa"}
	sort.Sort(ByNumVowels(words))
	fmt.Println(words)

}
