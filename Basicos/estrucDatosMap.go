package main

import "fmt"

func maps() {

	// Map

	myMap := make(map[string]int)
	myMap["Brais"] = 36
	myMap["Kevin"] = 26

	fmt.Println(myMap)
	fmt.Println(myMap["Brais"])

	myMap2 := map[string]int{"Brais": 36, "Kevin": 26}
	fmt.Println(myMap2)
	
}