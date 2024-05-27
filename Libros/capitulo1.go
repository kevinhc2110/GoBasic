package main

import (
	"fmt"
)

func capitulo1() {

	/*if len(os.Args) != 2 {
		os.Exit(1)
	}
	fmt.Println("It's over ", os.Args[1])*/

	// *go run main.go 9000 
	// Ejecutamos desde la terminal y le pasamos un argumento en este caso '9000'

	power := 1000
	fmt.Printf("default power is %d\n", power)
	
	// := se puede utilizar ya que al menos una de las variables es nueva
	// Aunque el tipo de la variable no puede ser cambiado pues ya fue declarada

	name, power := "Goku", 9000
	fmt.Printf("%s's power is over %d\n", name, power)


}

// *Ejecutar programa co fichero temporal
// go run main.go

// *Ejecutar programa viendo fichero temporal
// go run --work main.go

// *Compilar explícitamente (Generar binario)
// go build main.go

// El punto de entrada siempre es la función func main(){}

/*func variables(){

	// Declaración de variables

	// var power int = 9000
	// power := 9000

	// También se puede utilizar con funciones
	// power := getPower()
	
}*/

/*func getPower() int {
	return 9001
}*/

// Funciones

/*func log(message string) { // Función sin retorno
	}
	func add(a int, b int) int { // Función con retorno
	}
	func power(name string) (int, bool) { // Función con 2 retornos
	}
*/
	
/*func funciones(){
	
	value, exists := power("goku")
	if exists == false {
	// manejar esta situación de error
	/

	// O en el caso que solo estemos interesados uno de los retornos

	_, exists := power("goku")
	if exists == false {
	// manejar esta situación de error
	}
	
}
*/



