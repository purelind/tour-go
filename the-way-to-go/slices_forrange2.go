package main

import "fmt"

func main() {
	seasons := []string{"Spring", "Summer", "Autumn", "Winter"}
	for ix, season := range seasons {
		fmt.Printf("Season %d is: %s\n", ix, season)
	}

	for ix := range seasons {
		fmt.Printf("%d\n", ix)
	}

	var season string
	for _, season = range seasons {
		fmt.Printf("%s\n", season)
	}
}
