package datastructures

import (
	"fmt"
)



func ArrayCust(){
	var scores[3] float64
	students := [4] string{"Luffy", "Sanji", "Brook", "Zoro"}

	scores[0] = 89
	scores[1] = 36
	scores[2] = 76



	for index, name := range students{
		fmt.Println(index, name)
	}

	for index := 0; index < len(scores); index++{
		fmt.Println(scores[index])
	}



}


