package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"log"
	"os"
)

func readImageFile(path string) *os.File {
	file, err := os.Open(path)
	fmt.Printf("File type %T\n", file)
	if err != nil {
		return nil
	}
	defer file.Close()
	return file
}

func getImageSize(path string) (height int, width int) {
	file, err := os.Open(path)
	fmt.Printf("File type %T\n", file)
	if err != nil {
		log.Fatal("Error: ", err, "\n")
	}
	defer file.Close()
	config, _, err := image.DecodeConfig(file)
	if err != nil {
		log.Fatal("Error: ", err, "\n")
	}
	fmt.Println("Config: ", config)
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
	fmt.Println("Hello world!")
	// file := readImageFile("img.jpg")
	height, width := getImageSize("img.jpg")
	fmt.Println(height, width)
}
