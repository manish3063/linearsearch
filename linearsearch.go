package main

import "fmt"

func main() {

	array := []int{5, 3, 4, 2, 1, 6, 7, 8, 10, 9}

	var data int

	var found int

	found = 0

	data = 5
	var i int

	for i = 0; i < len(array); i++ {
		if array[i] == data {

			fmt.Printf("Found at index = %d\t", i+1)
			found = 1
			break

		}

	}
	if found == 0 {
		fmt.Println("data not found")

	}

}
