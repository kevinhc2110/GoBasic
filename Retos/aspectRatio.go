package main

import (
	"fmt"
	"image"
	_ "image/jpeg" // Importar el formato JPEG
	_ "image/png"  // Importar el formato PNG
	"net/http"
)

/*
Crea un programa que se encargue de calcular el aspect ratio de una
imagen a partir de una url.
- Url de ejemplo:
https://raw.githubusercontent.com/mouredevmouredev/master/mouredev_github_profile.png
- Por ratio hacemos referencia por ejemplo a los "16:9" de una
imagen de 1920*1080px.
*/

func aspectRatio() {

	// URL de la imagen

	url := "https://images.unsplash.com/photo-1606115915090-be18fea23ec7?q=80&w=1000&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8MXx8anBlZ3xlbnwwfHwwfHx8MA%3D%3D"

	// Realizar la petición HTTP para obtener la imagen

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error obteniendo la imagen: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// Decodificar la imagen

	img, _, err := image.Decode(resp.Body)
	if err != nil {
		fmt.Printf("Error decodificando la imagen: %v\n", err)
		return
	}

	// Obtener las dimensiones de la imagen

	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	// Calcular el aspect ratio

	aspectRatio := float64(width) / float64(height)

	// Imprimir el aspect ratio

	fmt.Printf("El aspect ratio de la imagen es: %.2f\n", aspectRatio)
}
