package main

// image: Proporciona funciones para manejar imágenes.
// image/color: Permite trabajar con colores (RGB, RGBA, etc.).
// image/jpeg, image/png, image/bmp: Para decodificar y codificar imágenes en diferentes.
// golang.org/x/image/bmp: Para trabajar con imágenes en formato BMP.
// path/filepath: Para manipulación de rutas de archivos.
// math: Necesaria para operaciones matemáticas como la raíz cuadrada.
// os: Para manejar archivos y el sistema de archivos.

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"math"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/image/bmp"
)

//* Problema 3 Filter

// abrirImagen abre una imagen desde el disco en formato BMP, JPEG o PNG
func abrirImagen(path string) (image.Image, string, error) {
	// Abre el archivo de imagen
	file, err := os.Open(path)
	if err != nil {
		return nil, "", err
	}
	defer file.Close()

	// Obtiene la extensión del archivo
	ext := strings.ToLower(filepath.Ext(path))

	var img image.Image
	// Dependiendo de la extensión, decodifica la imagen en el formato correcto
	switch ext {
	case ".jpg", ".jpeg":
		img, err = jpeg.Decode(file)
	case ".png":
		img, err = png.Decode(file)
	case ".bmp":
		img, err = bmp.Decode(file)
	default:
		// Si el formato no es soportado, devuelve un error
		return nil, "", fmt.Errorf("formato no soportado: %s", ext)
	}
	if err != nil {
		return nil, "", err
	}
	return img, ext, nil
}

// guardarImagen guarda la imagen en el disco en el formato original (JPEG, PNG o BMP)
func guardarImagen(path string, img image.Image, ext string) error {
	// Crea el archivo donde se guardará la imagen
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// Dependiendo de la extensión, codifica la imagen en el formato correcto
	switch ext {
	case ".jpg", ".jpeg":
		err = jpeg.Encode(file, img, nil)
	case ".png":
		err = png.Encode(file, img)
	case ".bmp":
		err = bmp.Encode(file, img)
	default:
		// Si el formato no es soportado, devuelve un error
		return fmt.Errorf("formato no soportado: %s", ext)
	}
	if err != nil {
		return err
	}
	return nil
}

// aplicarEscalaGrises convierte la imagen a escala de grises
func aplicarEscalaGrises(img image.Image) *image.RGBA {
	bounds := img.Bounds()
	grayImg := image.NewRGBA(bounds)

	// Recorre todos los píxeles de la imagen
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			// Obtiene los valores de los canales rojo, verde y azul
			r, g, b, _ := img.At(x, y).RGBA()
			// Calcula el promedio de los tres canales para generar el gris
			avg := uint8((r + g + b) / 3 >> 8)
			grayColor := color.RGBA{avg, avg, avg, 255}
			grayImg.Set(x, y, grayColor)
		}
	}
	return grayImg
}

// aplicarReflejo refleja la imagen horizontalmente
func aplicarReflejo(img image.Image) *image.RGBA {
	bounds := img.Bounds()
	reflectedImg := image.NewRGBA(bounds)

	// Recorre todos los píxeles de la imagen
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			// Copia el píxel en la posición reflejada
			originalPixel := img.At(x, y)
			reflectedImg.Set(bounds.Max.X-x-1, y, originalPixel)
		}
	}
	return reflectedImg
}

// aplicarDesenfoque aplica un desenfoque básico a la imagen
func aplicarDesenfoque(img image.Image) *image.RGBA {
	bounds := img.Bounds()
	blurredImg := image.NewRGBA(bounds)

	// Crea una copia de la imagen original para usar como referencia
	copyImg := image.NewRGBA(bounds)
	draw.Draw(copyImg, bounds, img, bounds.Min, draw.Src)

	// Recorre todos los píxeles de la imagen
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			var rSum, gSum, bSum, count uint32

			// Toma los píxeles vecinos (matriz de 3x3)
			for dy := -1; dy <= 1; dy++ {
				for dx := -1; dx <= 1; dx++ {
					px := x + dx
					py := y + dy

					// Asegura que los píxeles no estén fuera del límite de la imagen
					if px >= bounds.Min.X && px < bounds.Max.X && py >= bounds.Min.Y && py < bounds.Max.Y {
						r, g, b, _ := copyImg.At(px, py).RGBA()
						// Suma los valores de los canales RGB
						rSum += r
						gSum += g
						bSum += b
						count++
					}
				}
			}

			// Calcula el promedio para el efecto de desenfoque
			blurredImg.Set(x, y, color.RGBA{
				uint8(rSum / count >> 8),
				uint8(gSum / count >> 8),
				uint8(bSum / count >> 8),
				255,
			})
		}
	}
	return blurredImg
}

// aplicarDeteccionBordes aplica el filtro de detección de bordes (algoritmo Sobel)
func aplicarDeteccionBordes(img image.Image) *image.RGBA {
	bounds := img.Bounds()
	edgeImg := image.NewRGBA(bounds)

	// Crea una copia de la imagen original para usar como referencia
	copyImg := image.NewRGBA(bounds)
	draw.Draw(copyImg, bounds, img, bounds.Min, draw.Src)

	// Filtros Sobel (matrices para detectar cambios horizontales y verticales)
	gx := [3][3]int{
		{-1, 0, 1},
		{-2, 0, 2},
		{-1, 0, 1},
	}
	gy := [3][3]int{
		{-1, -2, -1},
		{0, 0, 0},
		{1, 2, 1},
	}

	// Recorre todos los píxeles de la imagen
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			var rX, gX, bX, rY, gY, bY int

			// Aplica las máscaras Sobel para calcular las derivadas en x y y
			for dy := -1; dy <= 1; dy++ {
				for dx := -1; dx <= 1; dx++ {
					px := x + dx
					py := y + dy

					if px >= bounds.Min.X && px < bounds.Max.X && py >= bounds.Min.Y && py < bounds.Max.Y {
						r, g, b, _ := copyImg.At(px, py).RGBA()
						weightX := gx[dy+1][dx+1]
						weightY := gy[dy+1][dx+1]

						rX += int(r>>8) * weightX
						gX += int(g>>8) * weightX
						bX += int(b>>8) * weightX

						rY += int(r>>8) * weightY
						gY += int(g>>8) * weightY
						bY += int(b>>8) * weightY
					}
				}
			}

			// Calcula la magnitud del gradiente
			r := uint8(math.Min(255, math.Sqrt(float64(rX*rX+rY*rY))))
			g := uint8(math.Min(255, math.Sqrt(float64(gX*gX+gY*gY))))
			b := uint8(math.Min(255, math.Sqrt(float64(bX*bX+bY*bY))))

			// Establece el píxel calculado en la nueva imagen
			edgeImg.Set(x, y, color.RGBA{r, g, b, 255})
		}
	}
	return edgeImg
}

func main() {
	// Abrir la imagen (se puede utilizar .jpg, .jpeg, .png o .bmp)
	img, ext, err := abrirImagen("./images/yard.bmp") // Cambia la ruta al archivo de imagen que quieras usar
	if err != nil {
		panic(err)
	}

	// Aplicar los diferentes filtros y guardar los resultados

	// Aplicar escala de grises y guardar la imagen
	grayImg := aplicarEscalaGrises(img)
	guardarImagen("grayscale"+ext, grayImg, ext)

	// Aplicar reflejo horizontal y guardar la imagen
	reflectedImg := aplicarReflejo(img)
	guardarImagen("reflected"+ext, reflectedImg, ext)

	// Aplicar desenfoque y guardar la imagen
	blurredImg := aplicarDesenfoque(img)
	guardarImagen("blurred"+ext, blurredImg, ext)

	// Aplicar detección de bordes y guardar la imagen
	edgesImg := aplicarDeteccionBordes(img)
	guardarImagen("edges"+ext, edgesImg, ext)

	fmt.Println("Procesamiento completado y imágenes guardadas.")
}
