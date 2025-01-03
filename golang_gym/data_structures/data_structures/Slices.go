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

func ReverseSlice(arr[]int){
	left := 0
	right := len(arr) - 1

	for left < right{
		arr[left], arr[right] = arr[right], arr[left]
		left += 1
		right -= 1
	}

	
}
func RemoveAllEven(arr[]int)[]int{
	var size int = (len(arr) + 1) / 2
	ans := make([]int, 0, size)

	for _ , curr_value := range arr{
		if curr_value % 2 == 1{
			ans = append(ans, curr_value)
		}
	}
	return ans

}
func RotateLeft(arr[]int, k int) []int{
	N := len(arr)
	ans := make([]int, 0 , N)
	k = k % N

	for index := k; index < N; index++{
		ans=append(ans, arr[index])

	}
	for index := range k{
		ans = append(ans, arr[index])
	}
	return ans

}
func RotateLeftInPlace(arr[]int , k int){
	N := len(arr)
	k = k % N

	for index := range k{
		arr[index], arr[N - (index + 1)] = arr[N - (index + 1)], arr[index]
	}

	left := 0
	right := N - k - 1
	// fmt.Println(arr)

	for left < right{
		arr[left], arr[right] = arr[right], arr[left]
		left += 1
		right -= 1
	}
	// fmt.Println(arr)
	left = N - k
	right = N - 1
	for left < right{
		arr[left], arr[right] = arr[right], arr[left]
		left += 1
		right -= 1
	}
	// fmt.Println(arr)

	
}