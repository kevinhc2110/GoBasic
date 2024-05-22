package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
DIFICULTAD EXTRA:
Desarrolla un programa de gestión de ventas que almacena sus datos en un
archivo .txt.
- Cada producto se guarda en una línea del archivo de la siguiente manera:
[nombre_producto], [cantidad_vendida], [precio].
- Siguiendo ese formato, y mediante terminal, debe permitir añadir, consultar,
actualizar, eliminar productos y salir.
- También debe poseer opciones para calcular la venta total y por producto.
- La opción salir borra el .txt.
*/

// Estructura del Producto
type Producto struct {
	Nombre          string
	CantidadVendida int
	Precio          float64
}

//Constante del archivo de ventas

const archivoVentas = "ventas.txt"

// Estas funciones se encargan de leer y escribir los productos en el archivo ventas.txt.
// Se utiliza para abrir el archivo especificado por os.Open()archivoVentas.
//Se comprueba si hay errores en la apertura mediante .if err != nil

func leerProductos() ([]Producto, error) {
	file, err := os.Open(archivoVentas)
	if err != nil {
		// Si se produce un error, se verifica si el error indica que el archivo no existe utilizando .os.IsNotExist()
		if os.IsNotExist(err) {
			return []Producto{}, nil
		}
		return nil, err
	}

	// Se utiliza defer file.Close() para garantizar que el archivo se cierre

	defer file.Close()

	//Se declara una variable de tipo productos[]Producto para almacenar los productos leídos del archivo.

	var productos []Producto
	// Se crea un nuevo escáner para leer el archivo línea a línea.bufio.NewScanner()
	scanner := bufio.NewScanner(file)
	// Se utiliza para leer cada línea del archivo.scanner.Scan()
	for scanner.Scan() {
		// Si la lectura es exitosa, se almacena la línea en la variable .linea
		linea := scanner.Text()
		// Se divide la línea en partes utilizando strings.Split() con el separador ", ".
		partes := strings.Split(linea, ", ")
		if len(partes) != 3 {
			continue
		}
		// cantidad se convierte a un entero utilizando strconv.Atoi()
		cantidad, _ := strconv.Atoi(partes[1])
		// precio se convierte a un número de coma flotante utilizando strconv.ParseFloat().
		precio, _ := strconv.ParseFloat(partes[2], 64)
		// Se crea un nuevo objeto Producto con los valores convertidos para Nombre, CantidadVendida y Precio.
		// Se agrega el nuevo objeto Producto a la lista productos.
		productos = append(productos, Producto{Nombre: partes[0], CantidadVendida: cantidad, Precio: precio})
	}
	// Se devuelve el error scanner.Err() si se produjo un error durante la lectura del archivo.
	return productos, scanner.Err()
}

func escribirProductos(productos []Producto) error {

	// Se utiliza para abrir o crear el archivo os.Create()archivoVentas.

	file, err := os.Create(archivoVentas)
	if err != nil {
		return err
	}
	defer file.Close()

	// Crea un escritor con búfer para escribir de forma eficiente en el archivo.

	writer := bufio.NewWriter(file)
	for _, producto := range productos {
		// utiliza fmt.Fprintf() para escribir el Nombre, CantidadVendida y Precio del producto en el escritor con búfer, separados por comas y formateando el precio con dos decimales.
		fmt.Fprintf(writer, "%s, %d, %.2f\n", producto.Nombre, producto.CantidadVendida, producto.Precio)
	}
	// Se llama a writer.Flush() para asegurar que todos los datos pendientes en el escritor con búfer se escriban en el archivo.
	return writer.Flush()

	// La función devuelve cualquier error encontrado durante el proceso de vaciado.
	
}

func añadirProducto(productos []Producto, producto Producto) []Producto {
	return append(productos, producto)
}

func buscarProducto(productos []Producto, nombre string) (Producto, int) {
	for i, producto := range productos {
		if strings.EqualFold(producto.Nombre, nombre) {
			return producto, i
		}
	}
	return Producto{}, -1
}

func actualizarProducto(productos []Producto, nombre string, nuevaCantidad int, nuevoPrecio float64) ([]Producto, bool) {
	for i, producto := range productos {
		if strings.EqualFold(producto.Nombre, nombre) {
			productos[i].CantidadVendida = nuevaCantidad
			productos[i].Precio = nuevoPrecio
			return productos, true
		}
	}
	return productos, false
}

func eliminarProducto(productos []Producto, nombre string) ([]Producto, bool) {
	for i, producto := range productos {
		if strings.EqualFold(producto.Nombre, nombre) {
			return append(productos[:i], productos[i+1:]...), true
		}
	}
	return productos, false
}

func calcularVentaTotal(productos []Producto) float64 {
	var total float64
	for _, producto := range productos {
		total += float64(producto.CantidadVendida) * producto.Precio
	}
	return total
}

func calcularVentaPorProducto(productos []Producto, nombre string) (float64, bool) {
	for _, producto := range productos {
		if strings.EqualFold(producto.Nombre, nombre) {
			return float64(producto.CantidadVendida) * producto.Precio, true
		}
	}
	return 0, false
}

func practica11()/*main*/ {
	reader := bufio.NewReader(os.Stdin)

	productos, err := leerProductos()
	if err != nil {
		fmt.Println("Error al leer los productos:", err)
		return
	}

	for {
		fmt.Println("\nOpciones:")
		fmt.Println("1. Añadir producto")
		fmt.Println("2. Consultar producto")
		fmt.Println("3. Actualizar producto")
		fmt.Println("4. Eliminar producto")
		fmt.Println("5. Calcular venta total")
		fmt.Println("6. Calcular venta por producto")
		fmt.Println("7. Salir")
		fmt.Print("Seleccione una opción: ")

		// lee caracteres del reader.ReadString('\n')reader hasta que encuentra un carácter de nueva línea (\n).
		opcionStr, _ := reader.ReadString('\n')
		// limina cualquier espacio en blanco inicial o final (espacios, tabuladores, nuevas líneas) de la cadena almacenada en . Esto asegura que la entrada del usuario, incluso si accidentalmente agregó espacios adicionales, se procese 
		opcionStr = strings.TrimSpace(opcionStr)
		// strconv.Atoi(opcionStr) intenta convertir la cadena en opcionStr (que se espera que sea un número ingresado por el usuario) en un valor entero.
		opcion, err := strconv.Atoi(opcionStr)
		if err != nil {
			fmt.Println("Opción no válida.")
			continue
		}

		switch opcion {
		case 1:
			fmt.Print("Nombre del producto: ")
			nombre, _ := reader.ReadString('\n')
			nombre = strings.TrimSpace(nombre)

			fmt.Print("Cantidad vendida: ")
			cantidadStr, _ := reader.ReadString('\n')
			cantidadStr = strings.TrimSpace(cantidadStr)
			cantidad, err := strconv.Atoi(cantidadStr)
			if err != nil {
				fmt.Println("Cantidad no válida.")
				continue
			}

			fmt.Print("Precio: ")
			precioStr, _ := reader.ReadString('\n')
			precioStr = strings.TrimSpace(precioStr)
			precio, err := strconv.ParseFloat(precioStr, 64)
			if err != nil {
				fmt.Println("Precio no válido.")
				continue
			}

			productos = añadirProducto(productos, Producto{Nombre: nombre, CantidadVendida: cantidad, Precio: precio})
			err = escribirProductos(productos)
			if err != nil {
				fmt.Println("Error al escribir el archivo:", err)
			} else {
				fmt.Println("Producto añadido exitosamente.")
			}

		case 2:
			fmt.Print("Nombre del producto: ")
			nombre, _ := reader.ReadString('\n')
			nombre = strings.TrimSpace(nombre)

			producto, _ := buscarProducto(productos, nombre)
			if (Producto{}) == producto {
				fmt.Println("Producto no encontrado.")
			} else {
				fmt.Printf("Producto encontrado: %s, Cantidad vendida: %d, Precio: %.2f\n", producto.Nombre, producto.CantidadVendida, producto.Precio)
			}

		case 3:
			fmt.Print("Nombre del producto: ")
			nombre, _ := reader.ReadString('\n')
			nombre = strings.TrimSpace(nombre)

			fmt.Print("Nueva cantidad vendida: ")
			cantidadStr, _ := reader.ReadString('\n')
			cantidadStr = strings.TrimSpace(cantidadStr)
			cantidad, err := strconv.Atoi(cantidadStr)
			if err != nil {
				fmt.Println("Cantidad no válida.")
				continue
			}

			fmt.Print("Nuevo precio: ")
			precioStr, _ := reader.ReadString('\n')
			precioStr = strings.TrimSpace(precioStr)
			precio, err := strconv.ParseFloat(precioStr, 64)
			if err != nil {
				fmt.Println("Precio no válido.")
				continue
			}

			productos, actualizado := actualizarProducto(productos, nombre, cantidad, precio)
			if actualizado {
				err = escribirProductos(productos)
				if err != nil {
					fmt.Println("Error al escribir el archivo:", err)
				} else {
					fmt.Println("Producto actualizado exitosamente.")
				}
			} else {
				fmt.Println("Producto no encontrado.")
			}

		case 4:
			fmt.Print("Nombre del producto: ")
			nombre, _ := reader.ReadString('\n')
			nombre = strings.TrimSpace(nombre)

			productos, eliminado := eliminarProducto(productos, nombre)
			if eliminado {
				err = escribirProductos(productos)
				if err != nil {
					fmt.Println("Error al escribir el archivo:", err)
				} else {
					fmt.Println("Producto eliminado exitosamente.")
				}
			} else {
				fmt.Println("Producto no encontrado.")
			}

		case 5:
			total := calcularVentaTotal(productos)
			fmt.Printf("Venta total: %.2f\n", total)

		case 6:
			fmt.Print("Nombre del producto: ")
			nombre, _ := reader.ReadString('\n')
			nombre = strings.TrimSpace(nombre)

			venta, encontrado := calcularVentaPorProducto(productos, nombre)
			if encontrado {
				fmt.Printf("Venta total para %s: %.2f\n", nombre, venta)
			} else {
				fmt.Println("Producto no encontrado.")
			}

		case 7:
			err := os.Remove(archivoVentas)
			if err != nil {
				fmt.Println("Error al borrar el archivo:", err)
			} else {
				fmt.Println("Archivo borrado exitosamente. Saliendo del programa.")
			}
			return

		default:
			fmt.Println("Opción no válida.")
		}
	}
}
