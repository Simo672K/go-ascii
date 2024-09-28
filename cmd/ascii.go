package cmd

import (
	"fmt"
	"os"

	"github.com/Simo672K/go-ascii/pkg/ascii"
	"github.com/Simo672K/go-ascii/pkg/process"
)

func ASCII() {
	file, err := os.Open("img.png")
	if err != nil {
		panic("Could not open image")
	}

	img := process.Image{
		ImgReader: file,
	}

	err = img.GetPixelArray()
	if err != nil {
		panic("Failed to get image pixel array")
	}

	img.ToGrayScale()
	generatedASCIIArt := ascii.GenerateASCII(img)

	fmt.Println(generatedASCIIArt)
}
