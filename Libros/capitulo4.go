// Capítulo 4 - Organización del Código e Interfaces

package main

// Paquetes

// Los nombres de los paquetes en Go siguen la estructura de directorios de tu workspace.
// cuando nombras un paquete mediante la palabra clave package, estás dando únicamente un valor, no una jerarquía completa (por ejemplo, “shopping” o “db”). Es necesario especificar la ruta completa a la hora de importar un paquete.

// En el directorio $GOPATH/src, se crea un directorio llamado shopping.
// Dentro de shopping, se crea un subdirectorio llamado db.
// Contenido del archivo db.go dentro de shopping/db:

/*package db

type Item struct {
    Price float64
}

func LoadItem(id int) *Item {
    return &Item{
        Price: 9.001,
    }
}*/

// Este archivo define un paquete llamado db que contiene una estructura Item y una función LoadItem que simula cargar un ítem desde una base de datos.

// Contenido del archivo pricecheck.go dentro de shopping:

/*package shopping

import "shopping/db"

func PriceCheck(itemId int) (float64, bool) {
    item := db.LoadItem(itemId)
    if item == nil {
        return 0, false
    }
    return item.Price, true
}
*/

// Este archivo define un paquete llamado shopping que importa el paquete db. Contiene una función PriceCheck que utiliza la función LoadItem del paquete db.

// Contenido del archivo main.go dentro de shopping/main:

/* package main

import (
    "fmt"
    "shopping"
)

func main() {
    fmt.Println(shopping.PriceCheck(4343))
}*/

// Este archivo define un programa ejecutable (main) que importa el paquete shopping y llama a la función PriceCheck para realizar una verificación de precios.
// Cuando importas shopping/db, en realidad estás importando $GOPATH/src/shopping/db, independientemente de tu ubicación dentro del paquete shopping.
// Para ejecutar este código desde la línea de comandos, puedes acceder al directorio shopping y usar go run para ejecutar el archivo main.go:

// cd $GOPATH/src/shopping
// go run main/main.go

// Importaciones Cíclicas

//Te irás encontrando con ellas a medida que vayas desarrollando software más complejo. Ocurren cuando el paquete A importa el paquete B a la vez que el paquete B importa el paquete A (ya sea directa o indirectamente a través de otro paquete). Esto es algo que el compilador no permite.

// Ejemplo
// Mueve la definición de Item de shopping/db/db.go a shopping/pricecheck.go. Tu fichero pricecheck.go debe quedar así:

/*package shopping
import (
31
"shopping/db"
)
type Item struct {
Price float64
}
func PriceCheck(itemId int) (float64, bool) {
item := db.LoadItem(itemId)
if item == nil {
return 0, false
}
return item.Price, true
}*/

// Si tratas de ejecutar el código te encontrarás con un par de errores de db/db.go indicando que Item no ha sido definido. Esto tiene sentido, ya que Item ha dejado de existir en el paquete db; ha sido movido al paquete shopping. Necesitamos modificar shopping/db/db.go a:

/*package db
import (
"shopping"
)
func LoadItem(id int) *shopping.Item {
return &shopping.Item{
Price: 9.001,
}
}*/

// Ahora, al tratar de ejecutar el código, obtenemos un intimidante error importación cíclica no permitida. Lo solucionaremos introduciendo otro paquete que contenga las estructuras compartidas. Tu estructura de directorios debería tener el siguiente aspecto:

/*$GOPATH/src
- shopping
pricecheck.go
- db
db.go
- models
item.go
- main
main.go*/

// pricecheck.go seguirá importando shopping/db, pero db.go importará shopping/models en lugar de shopping para así romper el ciclo. ya que hemos movido la estructura compartida Item a shopping/models/item.go, necesitamos cambiar shopping/db/db.go para referenciar la estructura Item desde el paquete models:

/*package db
import (
"shopping/models"
)
func LoadItem(id int) *models.Item {
return &models.Item{
Price: 9.001,
}
}*/

// A menudo necesitarás compartir algo más que sólamente models, así que quizá tengas otro directorio similar llamado utilities. La regla a tener en cuenta con respecto a los paquetes compartidos es que no deberían importar nada procedente del paquete shopping o de ninguno de sus subpaquetes. Dentro de un par de secciones veremos las interfaces, las cuales nos ayudarán a desenredar estos tipos de dependencias.

//*Visibilidad

// Go utiliza una regla muy sencilla para especificar qué tipos y qué funciones son visibles fuera del paquete. Si el nombre del tipo o de la función comienza con una letra mayúscula es visible. Si comienza con una letra minúscula no lo es.
// Esto también aplica a los campos de las estructuras. Si el nombre de un campo de una estructura comienza con una letra minúscula sólo el código que se encuentre dentro del mismo paquete será capaz de acceder a ella.

// Por ejemplo, si nuestro fichero items.go tiene una función con este aspecto:

/*func NewItem() *Item {
// ...
}*/

// podrá ser invocada mediante models.NewItem(). Pero si la función hubiese sido llamada newItem, no seríamos capaces de acceder a ella desde un paquete diferente.

// *Gestión de Paquetes

// El comando go que hemos estado utilizando con run y build también puede ser utilizado con get, que sirve para conseguir dependencias de terceros. go get funciona con varios protocolos aunque para este ejemplo, en el cual obtendremos una librería de Github, necesitarás tener git instalado en tu ordenador.
// Asumiendo que ya tienes git instalado, escribe lo siguiente desde un terminal:

/*go get github.com/mattn/go-sqlite3*/

// go get descarga los archivos y los almacena en tu workspace. Puedes comprobarlo echando un vistazo al directorio $GOPATH/src, verás que junto al proyecto que hemos creado antes encontrarás un directorio llamado github.com. Dentro de él verás un directorio mattn que contiene un directorio go-sqlite3.
// Acabamos de comentar cómo importar paquetes existentes en nuestro workspace. Para utilizar el paquete go-sqlite3 recién obtenido realizaremos el siguiente import:

/*import (
"github.com/mattn/go-sqlite3"
)
*/

//*Gestión de Dependencias

// go get guarda un par de ases bajo la manga. Si usamos go get en un proyecto se escanearán todos los ficheros buscando imports con librerías de terceros y las descargará. En cierta medida, nuestro propio código fuente se convierte en un Gemfile, un package.json, un pom.xml o un build.gradle.
// Si utilizas go get -u actualizarás todos los paquetes (también puedes actualizar un paquete específico usando go get -u NOMBRE_COMPLETO_DEL_PAQUETE)
// Puedes usar una herramienta de gestión de dependencias de terceros. Todavía son jóvenes pero las dos más prometedoras son goop y godep. Hay una lista más completa en la go-wiki.

//* Interfaces

//Los interfaces son tipos que especifican un contrato pero no contienen implementación. He aquí un ejemplo:

type Logger interface {
	Log(message string)
}

// Puede que te estés preguntando para qué puede servir esto. Los interfaces sirven para desacoplar tu código de implementaciones específicas. Por ejemplo, podríamos tener distintos tipos de loggers:

/*type SqlLogger struct { ... }
type ConsoleLogger struct { ... }
type FileLogger struct { ... }
*/

// ¿Cómo se utilizan? Exactamente igual que cualquier otro tipo. Puede ser el campo de una estructura:

/*type Server struct {
logger Logger
}*/

// o el parámetro de una función (o un valor de retorno):

/*func process(logger Logger) {
logger.Log("hello!")
}
*/

// Si tu estructura tiene una función llamada Log con un parámetro de tipo string y ningún valor de retorno, entonces puede ser utilizada como un Logger, lo que evita el tener que escribir tanto a la hora de usar interfaces:

/*type ConsoleLogger struct {}
func (l ConsoleLogger) Log(message string) {
fmt.Println(message)
}*/

// Esto facilita tener interfaces pequeños y muy enfocados en una tarea concreta. La librería estándar está llena de interfaces. Por ejemplo, el paquete io tiene varios de ellos, como io.Reader, io.Writer e io.Closer. Si escribes una función que espera un parámetro sobre el que invocar a Close() al terminar, entoces definitivamente deberías recibir io.Closer en lugar del tipo concreto que estés usando.
// Los interfaces también pueden participar en composiciones, de hecho, los interfaces pueden estar compuestos por otros interfaces. Por ejemplo, io.ReadCloser es un interfaz formado por los interfaces io.Reader e io.Closer.
