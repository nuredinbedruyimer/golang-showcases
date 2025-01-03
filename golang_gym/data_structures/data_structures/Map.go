package datastructures

import "fmt"


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