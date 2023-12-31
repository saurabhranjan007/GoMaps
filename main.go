package main

import "fmt"

func main() {
	fmt.Println("Hello from Go Maps ..")

	// Note📝: Maps in Go could be considered as Obj in JS or Dict in Python.
	// Maps are also stored in key-value parin, where keys could be of one data type and values could be of another data type.

	// Here this Map represents that all the keys will be of type string and similarly all the values be of type string 1️⃣
	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"white": "#0000",
	// 	"black": "#ffff",
	// }

	// fmt.Printf("%+v", colors)
	// fmt.Println(colors)

	// Another way of declaring and assigning Map 2️⃣
	// var color map[string]string
	// fmt.Println(color)

	// // ANother way of creating Map 3️⃣
	// color := make(map[string]string)
	// color["red"] = "#ff0000"
	// color["black"] = "#0000"
	// color["white"] = "#ffff"

	// // deleting values off of map
	// delete(color, "red")

	// fmt.Println(color)

	// Iterating ▶️ Over a Map
	color := map[string]string{
		"red":   "#ff0000",
		"white": "#0000",
		"black": "#ffff",
	}

	printMap(color)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, hex)
	}
}

// Why and where to use Map() and Struct()
