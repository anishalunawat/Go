package main

import "fmt"

func main() {
	/*var color map[string]string
	color := make(map[string]string)
	delete(color, "white")*/
	color := map[string]string{
		"red":   "RED",
		"black": "BLACK",
	}
	color["white"] = "WHITE"
	printMap(color)
}

func printMap(c map[string]string) {
	for color, blockColor := range c {
		fmt.Println(color, " ", blockColor)
	}
}
