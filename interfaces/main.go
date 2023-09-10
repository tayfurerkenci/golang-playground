package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreetings(eb)
	printGreetings(sb)
}

// Can use the same function name for different types via interfaces
func printGreetings(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	// VERY custom logic for generating an english greeting
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	// VERY custom logic for generating a spanish greeting
	return "Hola!"
}
