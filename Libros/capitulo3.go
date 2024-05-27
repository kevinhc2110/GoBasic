package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func capitulo3() {

	// 1. Crear una rebanada de enteros para almacenar puntuaciones

	scores := make([]int, 100)

	// 2. Llenar la rebanada con puntuaciones aleatorias entre 0 y 999

	for i := 0; i < 100; i++ {
		scores[i] = int(rand.Int31n(1000))
	}

	// 3. Ordenar las puntuaciones en orden ascendente (de menor a mayor)

	sort.Ints(scores)

	// 4. Crear una rebanada para almacenar las 5 peores puntuaciones (las más bajas)

	worst := make([]int, 5)

	// 5. Copiar los primeros 5 elementos (puntuaciones más bajas) de 'puntuaciones' a 'peores'

	copy(worst, scores[:5])
	
	fmt.Println(worst)
}

// *Arrays

// En Go, como en muchos otros lenguajes, los arrays tienen un tamaño fijo.

/*var scores [10]int
scores[0] = 339
*/

// El array anterior puede almacenar hasta 10 puntuaciones usando los índices que van desde scores[0] hasta scores [9].

// Es posible inicializar el array con valores:

/*scores := [4]int{9001, 9333, 212, 33*/

// Podemos utilizar len para obtener la longitud del array. range puede ser utilizado para iterar sobre él:

/*for index, value := range scores {
}
*/

// *Slices

//  Un slice es una estructura ligera que encapsula y representa una porción de un array.

/*scores := []int{1,4,293,4,9}*/

// Al contrario que con la declaración de arrays, nuestro slice no ha sido declarado indicando una longitud en los corchetes

// Vamos a ver otro modo de crear un slice, en este caso usando make, para entender por qué los dos son diferentes

/*scores := make([]int, 10)*/

// Usamos make en lugar de new porque la creación de un slice implica más cosas que simplemente reservar memoria (que es lo que hace new).
//  En concreto, tenemos que reservar la memoria para un array subyacente y, además, tenemos que inicializar el slice. En el ejemplo anterior inicializábamos un slice con una longitud de 10 y una capacidad de 10.

//  La longitud es el tamaño del slice y la capacidad es el tamaño del array subyacente. Utilizando make podemos especificar ambos por separado

/*scores := make([]int, 0, 10)
 */

// Esto crea un slice con longitud 0 y una capacidad de 10. (Si estás atento, quizá te habrás dado cuenta de que tanto make como len están sobrecargados.

// Vamos a ver algunos ejemplos con los que entender mejor la diferencia entre longitud y capacidad:

/*func main() {
scores := make([]int, 0, 10)
scores[5] = 9033
fmt.Println(scores)
}*/

// Nuestro primer ejemplo falla. ¿Por qué?, porque nuestro slice tiene una longitud de 0. Si, el array tiene una longitud de 10 elementos, pero necesitamos expandir nuestro slice explícitamente para poder acceder a esos elementos.

// forma de expandir un slice es a través de append:

/*func main() {
scores := make([]int, 0, 10)
scores = append(scores, 5)
fmt.Println(scores) // muestra [5]
}*/

// Pero esto cambia el propósito de nuestro código original. Hacer un append sobre un slice de tamaño 0 pondrá un valor en el primer elemento. Por la razón que sea, el código anterior quiere dar valor a la posición 5. Para conseguir esto podemos modificar nuestro slice:

/*func main() {
scores := make([]int, 0, 10)
scores = scores[0:6]
scores[5] = 9033
fmt.Println(scores)
}*/

// ¿Cuánto podemos redimensionar un slice? Hasta su capacidad, la cual en este caso es 10. Puede que estés pensando que eso no resuelve el problema de la longitud fija de los arrays. Esto convierte a append en algo especial. Si el array subyacente está lleno se creará uno más grande y se copiarán los valores sobre él.
// Este es el motivo por el cual, en el ejemplo anterior en el que usamos append, tenemos que reasignar el valor devuelto por append a nuestra variable scores: append puede haber creado un nuevo valor si el original se había quedado sin espacio.
// Si te dijera que Go hace crecer los arrays con un algoritmo que multiplica su tamaño por 2

/*func main() {
scores := make([]int, 0, 5)
c := cap(scores)
fmt.Println(c)
for i := 0; i < 25; i++ {
scores = append(scores, i)
// Si la capacidad ha cambiado,
// Go tiene que hacer crecer el array para volcar los datos
if cap(scores) != c {
c = cap(scores)
fmt.Println(c)
}
}
}
*/

// La capacidad inicial de scores es 5. Para poder almacenar 25 valores tiene que expandirse 3 veces para tener la capacidad de 10, 20 y finalmente 40.

/*func main() {
scores := make([]int, 5)
scores = append(scores, 9332)
fmt.Println(scores)
}
*/

// En este caso la salida será [0, 0, 0, 0, 0, 9332].

// *Inicializar un slice

/*names := []string{"leto", "jessica", "paul"}
checks := make([]bool, 10)
var names []string
scores := make([]int, 0, 20)*/

// La primera no debería tener mucha explicación: la deberías usar cuando sepas de antemano qué valores son los que debe haber en el array.
// El segundo es útil para escribir en índices específicos de un slice.

/*func extractPowers(saiyans []*Saiyans) []int {
powers := make([]int, len(saiyans))
for index, saiyan := range saiyans {
24
powers[index] = saiyan.Power
}
return powers
}
*/

// El tercero sirve para crear un slice vacío y se utiliza junto con append cuando el número de elementos es desconocido.
// La última versión nos permite indicar una capacidad inicial; es útil si tenemos una idea general de cuántos elementos vamos a necesitar.

// Puedes utilizar append incluso cuando sabes el tamaño. Es, de largo, una mera cuestión de preferencia:

/*func extractPowers(saiyans []*Saiyans) []int {
powers := make([]int, 0, len(saiyans))
for _, saiyan := range saiyans {
powers = append(powers, saiyan.Power)
}
return powers
}*/

// *Usar slices como wrappers de arrays es un concepto poderoso.

/*scores := []int{1,2,3,4,5}
slice := scores[2:4]
slice[0] = 999
fmt.Println(scores)*/

// La salida es [1, 2, 999, 4, 5

// Hay funciones que reciben la posición como un parámetro.
// strings.Index(haystack[5:], " ")
// Puedes ver en el ejemplo anterior que [X:] es un atajo que significa desde X hasta el fin mientras que [:X] es un atajo para desde el comienzo hasta X. Go, al contrario que otros lenguajes, no soporta valores negativos.

// Si queremos todos los valores de un slice excepto el último usamos
/*scores := []int{1, 2, 3, 4, 5}
scores = scores[:len(scores)-1]*/

// A continuación tienes una forma eficiente con la que borrar un valor de un slice sin ordenar:

/*func main() {
scores := []int{1, 2, 3, 4, 5}
scores = removeAtIndex(scores, 2)
fmt.Println(scores)
}
func removeAtIndex(source []int, index int) []int {
lastIndex := len(source) - 1
//intercambiamos el último valor con aquel que queremos eliminar
source[index], source[lastIndex] = source[lastIndex], source[index]
return source[:lastIndex]
}
*/

// *Maps

// Los mapas en Go son lo que en otros lenguajes se llaman tablas hash o diccionarios. Funcionan tal y como esperas: defines una clave y un valor y puedes recuperar, establecer y borrar valores en base a ella.
// Los mapas se crean con la función make al igual que los slices. Vamos a ver un ejemplo:

/*func main() {
lookup := make(map[string]int)
lookup["goku"] = 9001
power, exists := lookup["vegeta"]
// muestra 0, false
// 0 es el valor por defecto para un entero
fmt.Println(power, exists)
}
*/

// Usamos len para recuperar el número de claves. Para borrar un valor en base a su clave usamos delete:

/*// devuelve 1
total := len(lookup)
// no devuelve nada, puede ser invocado en una clave que no exista
delete(lookup, "goku")*/

// Los mapas crecen dinámicamente. Sin embargo, podemos pasarle un segundo argumento a make con el que indicar un tamaño inicial:

/*lookup := make(map[string]int, 100)*/

// Definir un tamaño inicial puede ayudar a mejorar el rendimiento, así que si tienes una idea sobre cuántas claves vas a tener, es bueno indicarlo al crear el mapa.

// Puedes definir del modo siguiente un mapa como campo en una estructura:

/*type Saiyan struct {
Name string
Friends map[string]*Saiyan
}*/

// Una forma de inicializar el mapa anterior es la siguiente:

/*goku := &Saiyan{
Name: "Goku",
Friends: make(map[string]*Saiyan),
}
goku.Friends["krillin"] = ... TODO: cargar o crear a Krillin*/

// Hay otra forma más de declarar e inicializar valores en Go

/*lookup := map[string]int{
"goku": 9001,
"gohan": 2044,
}*/

// Podemos iterar sobre un mapa usando un bucle for en combinación con la palabra clave range:

/*for key, value := range lookup {
...
}*/

// La iteración en mapas no está ordenada, de tal modo que cada iteración devolverá las parejas clave-valor en un orden aleatorio.

// * Punteros versus Valores

/*a := make([]Saiyan, 10)
//o
b := make([]*Saiyan, 10)*/

// La decisión sobre si definir un array de punteros o un array de valores depende de cómo utilices los valores individales, no de cómo utilices el array o el mapa.

