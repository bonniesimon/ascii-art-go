package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	"log"
	"os"
)

const BRIGHTNESS_CHARS string = "`^\",:;Il!i~+_-?][}{1)(|\\/tfjrxnuvczXYUJCLQ0OZmwqpdbkhao*#MW&8%B@$"

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

func getBrightnessMatrix(pixels [][]color.RGBA, size image.Point, filter func(r, g, b float64) float64) [][]uint8 {
	var bMatrix [][]uint8
	for i := 0; i < size.X; i++ {
		var row []uint8
		for j := 0; j < size.Y; j++ {
			r := float64(pixels[i][j].R)
			g := float64(pixels[i][j].G)
			b := float64(pixels[i][j].B)
			grey := uint8(filter(r, g, b))
			row = append(row, grey)
		}
		bMatrix = append(bMatrix, row)
	}

	return bMatrix
}

func averageFilter(r, g, b float64) float64 {
	return (r + g + b) / 3
}

func getBrightnessCharMatrix(bMatrix [][]uint8) [][]string {
	var brightnessValue uint8
	var bCharMatrix [][]string
	for i := 0; i < len(bMatrix)-1; i++ {
		var row []string
		for j := 0; j < len(bMatrix[i]); j++ {
			brightnessValue = bMatrix[i][j]
			bCharIndex := int(float64(brightnessValue) / 256.00 * float64(len(BRIGHTNESS_CHARS)))
			if bCharIndex == len(BRIGHTNESS_CHARS) {
				bCharIndex--
			}
			row = append(row, string(BRIGHTNESS_CHARS[bCharIndex]))
		}
		bCharMatrix = append(bCharMatrix, row)
	}
	return bCharMatrix
}

func main() {
	var imageFilePath string = "img.jpg"
	img := getImageFromFilePath(imageFilePath)

	var imageHeight, imageWidth int = img.Bounds().Max.Y, img.Bounds().Max.X
	fmt.Printf("Size: %d * %d\n", imageWidth, imageHeight)

	size := img.Bounds().Size()

	var pixels [][]color.RGBA = getPixelsFromImg(img)
	var bMatrix [][]uint8 = getBrightnessMatrix(pixels, size, averageFilter)

	str := getBrightnessCharMatrix(bMatrix)

	for i := 0; i < size.X; i++ {
		for j := 0; j < size.Y; j++ {
			fmt.Printf("%s%s%s", str[i][j], str[i][j], str[i][j])
		}
		fmt.Print("\n")
	}

}
