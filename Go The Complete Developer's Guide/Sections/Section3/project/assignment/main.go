package main

import "fmt"

//Prints list and if even says even, if odd says odd.
func main() {
	// Creates a slice of type int. 
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// Loop for each number in the slice numbers. 
	for number := range numbers {
		// If the number is devided by two and has a remainder of 0 then prints the number + "is even"
		if number % 2 == 0 {
			fmt.Println(number, "is even")
		// If not then prints the number + "is odd". 
		} else {
			fmt.Println(number, "is odd")
		}
	}
}
