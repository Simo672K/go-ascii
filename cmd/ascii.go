package cmd

import (
	"fmt"
	"os"

	"github.com/Simo672K/go-ascii/pkg/ascii"
	"github.com/Simo672K/go-ascii/pkg/cli"
	"github.com/Simo672K/go-ascii/pkg/process"
)

func ASCII() {
	imagePath, density := cli.GetArgs()
	fmt.Println(imagePath)
	file, err := os.Open(imagePath)
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

	generatedASCIIArt := ascii.GenerateASCII(img, density)
	fmt.Println(generatedASCIIArt)
}
