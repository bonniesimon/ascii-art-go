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
	fmt.Printf("%T\n", img.Bounds())
	fmt.Printf("Size: %d * %d\n", img.Bounds().Max.X, img.Bounds().Max.Y)
	fmt.Println(img.At(0, 0))
}
