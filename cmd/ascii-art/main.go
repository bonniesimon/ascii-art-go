package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"log"
	"os"
)

func getImageSize(path string) (height int, width int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Error: ", err, "\n")
	}
	defer file.Close()
	config, _, err := image.DecodeConfig(file)
	if err != nil {
		log.Fatal("Error: ", err, "\n")
	}
	return config.Height, config.Width
}

func getImageFromFilePath(path string) (image.Image, error) {
	file, err := os.Open(path)
	fmt.Printf("%T", file)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, x, err := image.Decode(file)
	fmt.Println("\nType: ", x)
	return img, err
}

func main() {
	height, width := getImageSize("img.jpg")
	fmt.Println(height, width)
}
