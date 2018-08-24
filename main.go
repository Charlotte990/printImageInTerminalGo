package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"

	"github.com/buger/goterm"
	"github.com/nfnt/resize"
)

func init() {
	// important or else At(), Bounds() functions will
	// caused memory pointer error!!
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
}

func main() {
	picture, err := getImage()
	if err != nil {
		log.Fatalf("cant decode image, err: %s", err)
	}

	smallerPicture := resizeToTerminalSize(picture)

	printMyPicture(smallerPicture)

}

func getImage() (image.Image, error) {
	myArgs := os.Args[1]
	fmt.Println(myArgs)
	convertFile := string(myArgs)
	//read image from file that already exists
	existingImageFile, err := os.Open(convertFile)
	if err != nil {
		fmt.Println("cant open image")
		return nil, err
	}
	defer existingImageFile.Close()
	//calling image.Decode() will give us the data and type of image it is as a string, not using type
	imageData, _, err := image.Decode(existingImageFile)
	if err != nil {
		log.Fatalf("cant decode image, err: %s", err)
	}
	return imageData, err
}

func resizeToTerminalSize(picture image.Image) image.Image {
	terminalHeight := uint(goterm.Height())
	terminalWidth := uint(goterm.Width())

	fmt.Println("Terminal height : ", terminalHeight)
	fmt.Println("Terminal width : ", terminalWidth)

	return resize.Resize(terminalWidth, terminalHeight, picture, resize.NearestNeighbor)
}

func printMyPicture(smallerPicture image.Image) {
	fmt.Println("PRINT MY PICTURE PART")
	myWidth := smallerPicture.Bounds().Size().X
	myHeight := smallerPicture.Bounds().Size().Y

	for yAxis := 0; yAxis < myHeight; yAxis++ {
		for xAxis := 0; xAxis < myWidth; xAxis++ {
			myColor := smallerPicture.At(xAxis, yAxis)
			r, g, b, _ := myColor.RGBA()
			averageColor := ((r + g + b) / 3) / 255
			if averageColor <= 35 {
				fmt.Print("█")
			} else if averageColor <= 120 {
				fmt.Print("_")
			} else if averageColor <= 255 {
				fmt.Print(".")
			} else {
				fmt.Print(" ")
			}
		}
	}
}
