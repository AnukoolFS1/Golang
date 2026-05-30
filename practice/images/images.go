package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func LoadImage() {
	file, err := os.Open("photo.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// img, err := png.Decode(file)
	// if err != nil {
	// 	panic(err)
	// }
}

func main() {
	img := image.NewRGBA(
		image.Rect(0, 0, 100, 100),
	)

	red := color.RGBA{255, 0, 0, 255}

	for y := 0; y < 100; y++ {
		for x := 0; x < 100; x++ {
			img.Set(x, y, red)
		}
	}

	file, _ := os.Create("red.png")
	defer file.Close()

	png.Encode(file, img)
}
