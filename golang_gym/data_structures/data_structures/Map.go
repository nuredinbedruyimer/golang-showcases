package datastructures

import (
	"fmt"
	"sort"
)


func MapCust(){
	nameFrequency := make(map[string]int)
	names := []string{
		"Luffy", 
		"Zoro", 
		"Sanji", 
		"Jimbei", 
		"Nami", 
		"Chopper", 
		"Franky", 
		"Brook", 
		"Robin",
		"Robin", "Brook", "Brook", "Luffy", "Luffy", "luffy", 
	}

	for _, name := range names{
		nameFrequency[name] += 1
	}

	for _, name := range names{
		nameFrequency[name] += 1
	}
	for name, freq := range nameFrequency{
		fmt.Println(name, freq)
	}
}

func FindFirstNonRepeatingChar(word string)rune{
	mapp := make(map[rune]int)
	for _, char := range word{
		mapp[char] += 1
	}
	for key, value := range mapp{
		if value == 1{
			return key
		}
	}
	return 'a'
}

func GroupAnagrams(words []string) [][]string {
	anagramGroups := make(map[string][]string)

	for _, word := range words {
		sorted := SortString(word)
		anagramGroups[sorted] = append(anagramGroups[sorted], word)
	}

	result := [][]string{}
	for _, group := range anagramGroups {
		result = append(result, group)
	}
	return result
}

func SortString(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}
