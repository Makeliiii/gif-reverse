package main

import (
	"image/gif"
	"log"
	"os"
)


func main() {
	readPath := "gif/booba.gif"
	writePath := "gif/booba.jpeg"
	reader, err := os.Open(readPath)
	if err != nil {
		log.Fatal(err)
	}

	writer, errrr := os.Create(writePath)
	if errrr != nil {
		log.Fatal(errrr)
	}

	image, error := gif.Decode(reader)
	if error != nil {
		log.Fatal(error)
	}

	gif.Encode(writer, image, &gif.Options{NumColors: 256})
}