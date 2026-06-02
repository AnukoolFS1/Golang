package main

// --------the interface of image

// type Image interface{
// 	ColorModel() color.Model
//     Bounds() Rectangle
//     At(x, y int) color.Color
// }

// 1. Bounds(): Returns a Rectangle defining the image’s dimensions (e.g., from (0,0) to (width, height)).
// 2. At(x, y): Returns the color.Color at a specific pixel coordinate.
// 3. ColorModel(): Tells you how the image stores color data (e.g., RGB, RGBA, Grayscale).


import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	// 1. Define the dimensions
	width, height := 1200, 1200
	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	// 2. Create a blank RGBA image canvas
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// 3. Fill the canvas with colors
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			// Let's make a cool gradient!
			pixelColor := color.RGBA{uint8(float64(x) / float64(width) * 255), uint8(float64(y) / float64(height) * 255), 80, 255}
			img.Set(x, y, pixelColor)
		}
	}

	// 4. Save the image to a file
	f, _ := os.Create("gradient.png")
	defer f.Close()
	png.Encode(f, img)
}