package main

import (
	"fmt"
	"image"
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

func averageFilter(r, g, b uint32) int {
	return (int(r) + int(g) + int(b)) / 3
}

func rgba2rgb(r, g, b, a uint32) {
	red := a * r
	green := a * g
	blue := a * b
	fmt.Println(red, " ", green, " ", blue)
}

func main() {
	var imageFilePath string = "img.jpg"
	img := getImageFromFilePath(imageFilePath)

	var imageHeight, imageWidth int = img.Bounds().Max.Y, img.Bounds().Max.X
	fmt.Printf("Size: %d * %d\n", imageWidth, imageHeight)

	// var brightnessMatrix [][]int

	// Get Pixel values
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			fmt.Printf("%d\n", img.At(x, y))
			r, g, b, a := img.At(x, y).RGBA()
			rgba2rgb(r, g, b, a)
			fmt.Println(((r >> 8) & 0xFF))
			fmt.Printf("r: %T\n", r)
			fmt.Printf("r: %d g: %d b: %d", r, g, b)
			// fmt.Println(int(r), " ", int(g), " ", int(b))
			fmt.Println(averageFilter(r, g, b))
			break
		}
		break
	}
}
