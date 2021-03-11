package main

import (
	"fmt"
	"image/gif"
	"log"
	"os"
)


func main() {
	readPath := "gif/booba.gif"
	writePath := "gif/boobaRev.gif"
	reader, err := os.Open(readPath)
	if err != nil {
		log.Fatal(err)
	}

	writer, errrr := os.Create(writePath)
	if errrr != nil {
		log.Fatal(errrr)
	}

	images, error := gif.DecodeAll(reader)
	if error != nil {
		log.Fatal(error)
	}

	//revArr := make([]image.PalettedImage, 0, cap(images.Image))

	//fmt.Println("Before loop: ", images)

	newGif := gif.GIF{}
	newGif.LoopCount = 0

	for i := len(images.Image) - 1; i > -1; i-- {
		newGif.Image = append(newGif.Image, images.Image[i])
		newGif.Delay = append(newGif.Delay, images.Delay[i])
		newGif.Disposal = append(newGif.Disposal, images.Disposal[i])
	}

	fmt.Println("After loop:", newGif)

	gif.EncodeAll(writer, &newGif)
}