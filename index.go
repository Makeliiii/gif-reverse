package main

import (
	"fmt"
	"image/gif"
	"os"
)

func main() {
	readPath := "gif/booba.gif"
	writePath := "gif/boobaRev.gif"
	
	// ignore errors like any decent programmer
	reader, _ := os.Open(readPath)
	writer, _ := os.Create(writePath)
	images, _ := gif.DecodeAll(reader)

	// new gif yeya
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