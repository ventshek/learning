package main

import "fmt"

//Prints out an int array of numbers from 0 to 10
//Prints list and if even says even, if odd says odd.
func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for number := range numbers {
		if number % 2 == 0 {
			fmt.Println(number, "is even")
		} else {
			fmt.Println(number, "is odd")
		}
	}
}
