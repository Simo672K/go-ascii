package cmd

import (
	"fmt"
	"os"

	"github.com/Simo672K/go-ascii/pkg/process"
)

func ASCII() {
	file, err := os.Open("golang.png")
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
	fmt.Println(img.GrayScaledPixelArray)
}
