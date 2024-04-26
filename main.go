package main

import (
	"fmt"
)

func main() {
	numbers := [6]int{1, 15, 6, 7, 9, 5}
	for i := 0; i < (len(numbers) - 1); i++ {
		for j := 0; j < ((len(numbers) - 1) - i); j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}
	fmt.Println("")
	fmt.Printf("Sorted list:\t%v\n", numbers)
}
