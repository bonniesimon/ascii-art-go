package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	"log"
	"os"
)

func getImageFromFilePath(path string) image.Image {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Error: ", err, "\n")
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	return img
}

func getPixelsFromImg(img image.Image) [][]color.RGBA {
	var pixels [][]color.RGBA

	for i := 0; i < img.Bounds().Size().X; i++ {
		var y []color.RGBA
		for j := 0; j < img.Bounds().Size().Y; j++ {
			var c color.RGBA = color.RGBAModel.Convert(img.At(i, j)).(color.RGBA)
			y = append(y, c)
		}
		pixels = append(pixels, y)
	}

	return pixels
}

func averageFilter(r, g, b uint32) int {
	return (int(r) + int(g) + int(b)) / 3
}

func main() {
	var imageFilePath string = "img.jpg"
	img := getImageFromFilePath(imageFilePath)

	var imageHeight, imageWidth int = img.Bounds().Max.Y, img.Bounds().Max.X
	fmt.Printf("Size: %d * %d\n", imageWidth, imageHeight)

	size := img.Bounds().Size()

	var pixels [][]color.RGBA = getPixelsFromImg(img)

	for i := 0; i < size.X; i++ {
		for j := 0; j < size.Y; j++ {
			fmt.Printf("R: %d G: %d B: %d\n", pixels[i][j].R, pixels[i][j].G, pixels[i][j].B)
		}
	}

}
