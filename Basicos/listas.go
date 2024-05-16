package main

import (
	"container/list"
	"fmt"
)

func listas() {

	// Lista

	myList := list.New()
	myList.PushBack(1)
	myList.PushBack(2)
	myList.PushBack(3)

	fmt.Println(myList)              //Representación de punteros
	fmt.Println(myList.Back().Value) // Representa el ultimo valor de la lista
}