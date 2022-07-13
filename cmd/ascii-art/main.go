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

func main() {
	var imageFilePath string = "img.jpg"
	img := getImageFromFilePath(imageFilePath)
	fmt.Printf("Size: %d * %d\n", img.Bounds().Max.X, img.Bounds().Max.Y)

	// Get Pixel values
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			fmt.Println(img.At(x, y))
		}
	}
}
