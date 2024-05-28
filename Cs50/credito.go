package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func credito() {

var cardNumber string

fmt.Println("Ingrese un numero de tarjeta")
fmt.Scanln(&cardNumber)	

if luhnCheck(cardNumber) {
	cardType := getCardType(cardNumber)
	fmt.Printf("La tarjeta es válida y es de tipo %s\n", cardType)
} else {
	fmt.Println("La tarjeta no es válida")
}


}

func luhnCheck(cardNumber string) bool {

	sum := 0
	isSecond := false

	// Recorremos los dígitos del número de tarjeta desde el final

	for i := len(cardNumber) - 1; i >= 0; i-- {
		digit, _ := strconv.Atoi(string(cardNumber[i]))

		// Multiplicamos por 2 cada segundo dígito empezando desde el final

		if isSecond {
			digit *= 2

			// Si el producto es mayor a 9, restamos 9

			if digit > 9 {
				digit -= 9
			}
		}

		// Sumamos el dígito al total

		sum += digit

		// Alternamos entre dígitos para multiplicar por 2

		isSecond = !isSecond
	}

	// La tarjeta es válida si la suma total es múltiplo de 10

	return sum%10 == 0
}

func getCardType(cardNumber string) string {

	// Uso de expresiones regulares

	visaPattern := regexp.MustCompile("^4[0-9]{12}(?:[0-9]{3})?$")
	mastercardPattern := regexp.MustCompile("^5[1-5][0-9]{14}$")
	amexPattern := regexp.MustCompile("^3[47][0-9]{13}$")

	if visaPattern.MatchString(cardNumber) {
		return "Visa"
	} else if mastercardPattern.MatchString(cardNumber) {
		return "Mastercard"
	} else if amexPattern.MatchString(cardNumber) {
		return "American Express"
	}

	return "Desconocida"
}
