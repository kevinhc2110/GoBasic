package main

import "fmt"


type Saiyan struct {
	Name  string
	Power int
}

func (s *Saiyan) Super() {
	s.Power += 10000
}

// *Podemos asociar un método con una estructura
// En el código anterior podemos decir que el tipo *Saiyan es el receptor del método Super. Podemos llamar a Super de este modo

func capitulo2() {
	goku := &Saiyan{"Goku", 9001}
	goku.Super()
	fmt.Println(goku.Power) // mostrará 19001
}

// *Estructuras

/*type Saiyan struct {
Name string
Power int
}
*/

// *Declaras e inicializar estructura

/*goku := Saiyan{
Name: "Goku",
Power: 9000,
*/

// Nota: La última , del ejemplo anterior es obligatoria
// Tambien podemos

// No tenemos por qué dar valor a todos los campos. Ambas declaraciones son válidas
// goku := Saiyan{}
// o
// goku := Saiyan{Name: "Goku"}
// goku.Power = 9000

// Puedes omitir el nombre de los campos y basarte únicamente en el orden en el cual éstos están declarados
// goku := Saiyan{"Goku", 9000}

// *Trabajar con punteros

/*func main() {
goku := Saiyan{"Goku", 9000}
Super(goku)
fmt.Println(goku.Power)
}
func Super(s Saiyan) {
s.Power += 10000
}
*/
// La respuesta es 9000, no 19000. ¿Por qué? Porque Super hizo cambios en una copia de nuestro goku original y, por eso, los cambios realizados en Super no se vieron reflejados en quien hizo la invocación.
// Para hacer que todo funcione como probablemente esperabas necesitamos pasar un puntero a nuestra variable:

/*func main() {
goku := &Saiyan{"Goku", 9000}
Super(goku)
fmt.Println(goku.Power)
}
func Super(s *Saiyan) {
s.Power += 10000
}
*/

// Hemos hecho dos cambios. El primero es la utilización del operador & para recuperar la dirección de nuestra variable
// A continuación hemos modificado el tipo de parámetro que espera recibir Super. Antes esperaba un valor de tipo Saiyan pero ahora espera una dirección de tipo *Saiyan, donde *X significa puntero a valor de tipo X. Hay una relación obvia entre los tipos Saiyan y *Saiyan, pero son tipos distintos.

// *Constructores

// Las estructuras no tienen constructores. Sin embargo, puedes crear una función que devuelva una instancia del tipo deseado en su lugar (como un patrón factory):

/*func NewSaiyan(name string, power int) *Saiyan {
	return &Saiyan{
	Name: name,
	Power: power,
	}
}*/

// Nuestra factoría no tiene por qué devolver un puntero; este trozo de código es absolutamente válido:

/*func NewSaiyan(name string, power int) Saiyan {
	return Saiyan{
	Name: name,
	Power: power,
	}
}*/

// *New

// A pesar de la falta de constructores, Go tiene la función predefinida new que se utiliza para reservar la memoria que un tipo concreto necesita. El resultado de new(X) es el mismo que &X{}:

/*goku := new(Saiyan)
// es lo mismo que
goku := &Saiyan{}*/

// Cuál uses depende de ti, pero verás que la mayoría de la gente prefiere el segundo cuando hay campos que inicializar, ya que suele ser más fácil de leer:

/*goku := new(Saiyan)
goku.name = "goku"
goku.power = 9001
//vs
goku := &Saiyan {
name: "goku",
power: 9000,
}
*/

// *Campos de una Estructura

// Saiyan, en el ejemplo que hemos visto hasta ahora, tiene dos campos: Name y Power, de tipos string e int, respectivamente.
//  Podemos hacer crecer nuestra definición de Saiyan:

/*type Saiyan struct {
Name string
Power int
Father *Saiyan
}
*/

// la cual puede ser inicializada así:

/*gohan := &Saiyan{
Name: "Gohan",
Power: 1000,
Father: &Saiyan {
Name: "Goku",
Power: 9001,
Father: nil,
},
}
*/

// *Composición

// Go incluye el concepto de composición, que consiste en incluir una estructura dentro de otra. En algunos lenguajes esto se conoce como trait o mixin.

/*type Person struct {
Name string
}
func (p *Person) Introduce() {
fmt.Printf("Hi, I'm %s\n", p.Name)
}
type Saiyan struct {
*Person
Power int
}
// cómo usarlo:
goku := &Saiyan{
Person: &Person{"Goku"},
Power: 9001,
}
goku.Introduce()
*/

// La estructura Saiyan posee un campo de tipo *Person. Dado que no le hemos puesto explícitamente un nombre podemos implícitamente acceder a los campos y funciones del tipo compuesto.

/*goku := &Saiyan{
Person: &Person{"Goku"},
}
fmt.Println(goku.Name)
fmt.Println(goku.Person.Name)
*/

// Ambas expresiones son perfectamente válidas

// *Sobrecarga

// Aunque la sobrecarga no sea un concepto específico de las estructuras merece la pena mencionarlo. Go, sencillamente, no soporta sobrecarga. Por este motivo verás (y escribirás) muchas funciones con la firma Load, LoadById, LoadByName, etc.

// Por ejemplo, nuestra estructura Saiyan puede tener su propia función Introduce:

/*func (s *Saiyan) Introduce() {
fmt.Printf("Hi, I'm %s. Ya!\n", s.Name)
}
*/

// La versión original siempre estará accesible a través de s.Person.Introduce().

	
	
