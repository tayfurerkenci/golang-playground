package main

import "fmt"

func main() {
	// Creating maps
	// Method 1
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	// Method 2
	// var colors map[string]string

	// Method 3
	// colors := make(map[string]string)

	// Adding values to maps
	// colors["white"] = "#ffffff"

	// Deleting values from maps
	// delete(colors, "white")

	printMap(colors)
}

func printMap(c map[string]string) {
	fmt.Println(c)
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
