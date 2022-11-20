package main

import "fmt"

// Our program has a new type called bot.
type bot interface {
	// If you have a type in the program with a function called 'getGreeting' and you return a string then you are now an honorary member of type 'bot'.
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(sb) 
	printGreeting(eb)


}

// Now that you're also an honorary member of type 'bot', you can now call this function called 'printGreetint'
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())

}

func (eb englishBot) getGreeting() string {
	// VERY custom logic for generating an english greeting 
	return "Hi there!"
}


func (sb spanishBot) getGreeting() string {
	return "Hola!"
}

