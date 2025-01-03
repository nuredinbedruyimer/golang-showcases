package main

import (
	datastructures "github.com/nuredinbedru/data_structures/data_structures"
)

func main(){

	arr := []int{1, 2, 3, 4, 5}
	k := 7
	datastructures.RotateLeftInPlace(arr, k)
	print(arr)
	

	
}