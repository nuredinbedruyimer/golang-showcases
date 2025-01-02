package datastructures

import "fmt"



func SliceCust(){

	students := []string{"Luffy", "Sanji", "Zoro", "Brook"}
	students = append(students, "Nami")
	students = append(students, "Robin")
	students = append(students, "Chooper")

	for index,  name := range students{
		fmt.Println(index, name)
	}
}