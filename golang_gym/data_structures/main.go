package main

import (
	"fmt"

	datastructures "github.com/nuredinbedru/data_structures/data_structures"
)

func main(){

	words := []string{"bat", "tab", "cat", "act", "tac"}
	fmt.Println(    datastructures.GroupAnagrams(words))
	fmt.Println(words)
	

	
}